package event

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"github.com/lin-snow/ech0/internal/agent"
	commonModel "github.com/lin-snow/ech0/internal/model/common"
	echoModel "github.com/lin-snow/ech0/internal/model/echo"
	settingModel "github.com/lin-snow/ech0/internal/model/setting"
	"github.com/lin-snow/ech0/internal/persona"
	echoRepository "github.com/lin-snow/ech0/internal/repository/echo"
	inboxRepository "github.com/lin-snow/ech0/internal/repository/inbox"
	keyvalue "github.com/lin-snow/ech0/internal/repository/keyvalue"
	todoRepository "github.com/lin-snow/ech0/internal/repository/todo"
	userRepository "github.com/lin-snow/ech0/internal/repository/user"
)

type AgentProcessor struct {
	echoRepo     echoRepository.EchoRepositoryInterface
	todoRepo     todoRepository.TodoRepositoryInterface
	userRepo     userRepository.UserRepositoryInterface
	keyvalueRepo keyvalue.KeyValueRepositoryInterface
	inboxRepo    inboxRepository.InboxRepositoryInterface
}

func NewAgentProcessor(
	echoRepo echoRepository.EchoRepositoryInterface,
	todoRepo todoRepository.TodoRepositoryInterface,
	userRepo userRepository.UserRepositoryInterface,
	keyvalueRepo keyvalue.KeyValueRepositoryInterface,
	inboxRepo inboxRepository.InboxRepositoryInterface,
) *AgentProcessor {
	return &AgentProcessor{
		echoRepo:     echoRepo,
		todoRepo:     todoRepo,
		userRepo:     userRepo,
		keyvalueRepo: keyvalueRepo,
		inboxRepo:    inboxRepo,
	}
}

func (ap *AgentProcessor) Handle(ctx context.Context, e *Event) error {
	// 获取 Agent 设置
	var agentSetting settingModel.AgentSetting
	if agentSettingStr, err := ap.keyvalueRepo.GetKeyValue(commonModel.AgentSettingKey); err == nil {
		json.Unmarshal([]byte(agentSettingStr.(string)), &agentSetting)
	}

	// 清理生成内容的缓存
	ap.clearCache()

	// 更新平行人格
	ap.updatePersona(&agentSetting, e)

	return nil
}

func (ap *AgentProcessor) clearCache() error {
	// 删除 AGENT_GEN_RECENT 缓存
	ap.keyvalueRepo.DeleteKeyValue(context.Background(), string(agent.GEN_RECENT))

	return nil
}

func (ap *AgentProcessor) updatePersona(setting *settingModel.AgentSetting, e *Event) error {
	// 配置并开启了 Agent 才能更新人格
	if setting == nil || !setting.Enable {
		return nil
	}

	// 取出 Echo
	payload := e.Payload[EventPayloadEcho]
	echo, ok := payload.(echoModel.Echo)
	if !ok {
		return nil
	}

	// 取出当前人格
	var p persona.Persona
	if personaStr, err := ap.keyvalueRepo.GetKeyValue(persona.PersonaKey); err == nil {
		json.Unmarshal([]byte(personaStr.(string)), &p)
	} else {
		// 如果没有找到对应的人格，初始化一个默认人格
		now := time.Now().Unix()
		p = persona.Persona{
			Name:         "Persona",
			Description:  "parallel personality",
			Style:        []persona.Feature{},
			Mood:         []persona.Feature{},
			Topics:       []persona.Feature{},
			Expression:   []persona.Feature{},
			Independence: 0.5,
			CreatedAt:    now,
			UpdatedAt:    now,
			LastActive:   now,
		}
	}

	// 随机获取一个维度进行更新
	dim := p.WhatDimensionToUpdate()
	features := p.GetDimensionFeatures(dim)

	// 构建大模型输入
	template := prompt.FromMessages(schema.FString,
		schema.UserMessage(`
你是一套“人格特征更新器”，你的任务是根据输入内容更新某个人格维度的特征。  
你必须严格遵守以下规则：

任务：
根据给定的【维度】【已有特征】【用户近期行为（Echo）】生成新的特征列表（Feature 数组）。

输出格式（必须严格遵守）：
[
{"name": "中文特征名", "weight": 0.xx},
{"name": "中文特征名", "weight": 0.xx}
]

规则要求：
1. 所有特征名称必须是中文的、不带标点、简短的词语或短语。
2. 特征必须从属于指定维度：“{dimension}”。
3. weight 必须是 0~1 的浮点数。
4. 最终输出必须是合法 JSON，禁止输出任何解释性文本。

维度说明：
style（风格维度）：行为方式、说话风格，如：温和、犀利、冷静、机敏。
mood（情绪维度）：情绪状态，如：愉快、紧张、轻松、烦躁。
topics（兴趣主题维度）：偏好讨论的主题，如：科技、生活、哲学、编程。
expression（表达偏好维度）：表达方式，如：简洁表达、比喻表达、故事表达。

当前维度：
{dimension}

当前特征列表（可能为空）：
{features}

用户最近行为（Echo）：
{echo}

请基于以上内容生成新的完整特征列表（替换旧列表，而不是增量），并直接输出 JSON 数组。
不允许输出任何额外文字。
`),
	)
	featuresJSON, _ := json.Marshal(features)
	vars := map[string]any{
		"dimension": string(dim),
		"features":  string(featuresJSON),
		"echo":      echo.Content,
	}
	in, err := template.Format(context.Background(), vars)
	if err != nil {
		return err
	}

	// 调用大模型 根据 Echo 内容和 选中的维度，生成新的特征，得到未校验的Feature列表
	out, err := agent.Generate(context.Background(), *setting, in, false)
	if err != nil {
		return err
	}

	// 轻度 JSON 修复
	out = strings.TrimSpace(out)
	l := strings.Index(out, "[")
	r := strings.LastIndex(out, "]")
	if l >= 0 && r >= 0 {
		out = out[l : r+1]
	}

	// 解析大模型输出，得到新的特征列表
	var newFeatures []persona.Feature
	if err := json.Unmarshal([]byte(out), &newFeatures); err != nil {
		return err
	}

	// 执行更新
	p.UpdateDimension(dim, newFeatures)
	p.UpdatedAt = time.Now().Unix()

	// 保存更新后的人格
	personaBytes, err := json.Marshal(p)
	if err != nil {
		return err
	}
	if err := ap.keyvalueRepo.AddOrUpdateKeyValue(context.Background(), persona.PersonaKey, string(personaBytes)); err != nil {
		return err
	}

	return nil
}
