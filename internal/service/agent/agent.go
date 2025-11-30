package service

import (
	"context"
	"errors"

	"github.com/cloudwego/eino-ext/components/model/claude"
	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino-ext/components/model/gemini"
	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino-ext/components/model/qwen"
	"github.com/cloudwego/eino/schema"
	model "github.com/lin-snow/ech0/internal/model/setting"
	"google.golang.org/genai"

	authModel "github.com/lin-snow/ech0/internal/model/auth"
	commonModel "github.com/lin-snow/ech0/internal/model/common"
	echoService "github.com/lin-snow/ech0/internal/service/echo"
	settingService "github.com/lin-snow/ech0/internal/service/setting"
	todoService "github.com/lin-snow/ech0/internal/service/todo"
)

type AgentService struct {
	settingService settingService.SettingServiceInterface
	echoService    echoService.EchoServiceInterface
	todoService    todoService.TodoServiceInterface
}

func NewAgentService(
	settingService settingService.SettingServiceInterface,
	echoService echoService.EchoServiceInterface,
	todoService todoService.TodoServiceInterface,
) AgentServiceInterface {
	return &AgentService{
		settingService: settingService,
		echoService:    echoService,
		todoService:    todoService,
	}
}

func (agentService *AgentService) GetRecent(ctx context.Context) (string, error) {
	echos, err := agentService.echoService.GetEchosByPage(authModel.NO_USER_LOGINED, commonModel.PageQueryDto{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		return "", err
	}

	var memos []*schema.Message
	for _, echo := range echos.Items {
		memos = append(memos, &schema.Message{
			Role:    schema.Assistant,
			Content: echo.Content,
		})
	}

	in := []*schema.Message{
		{
			Role:    schema.System,
			Content: "你是一个热心的个人助理，帮助用户回顾最近的活动。",
		},
		{
			Role:    schema.User,
			Content: "请根据我最近的活动进行总结。",
		},
	}

	in = append(in, memos...)

	output, err := agentService.Generate(ctx, in)
	if err != nil {
		return "", err
	}

	return output, nil
}

func (agentService *AgentService) Generate(ctx context.Context, in []*schema.Message) (string, error) {
	var setting model.AgentSetting
	if err := agentService.settingService.GetAgentSettings(&setting); err != nil {
		return "", errors.New(commonModel.AGENT_SETTING_NOT_FOUND)
	}

	if !setting.Enable {
		return "", errors.New(commonModel.AGENT_NOT_ENABLED)
	}
	if setting.Model == "" {
		return "", errors.New(commonModel.AGENT_MODEL_MISSING)
	}
	if setting.Provider == "" {
		return "", errors.New(commonModel.AGENT_PROVIDER_NOT_FOUND)
	}
	if setting.ApiKey == "" {
		return "", errors.New(commonModel.AGENT_API_KEY_MISSING)
	}

	baseURL := ""
	if setting.BaseURL != "" {
		baseURL = setting.BaseURL
	}

	apiKey := setting.ApiKey
	model := setting.Model

	var resp *schema.Message
	var genErr error

	// 选择服务提供商
	switch setting.Provider {
	case string(commonModel.OpenAI):
		cm, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
			APIKey:  apiKey,
			Model:   model,
			BaseURL: baseURL,
		})

		if err != nil {
			return "", err
		}

		resp, genErr = cm.Generate(ctx, in)

	case string(commonModel.Anthropic):
		var baseURLPtr *string = nil
		if len(baseURL) > 0 {
			baseURLPtr = &baseURL
		}

		cm, err := claude.NewChatModel(ctx, &claude.Config{
			APIKey:  apiKey,
			Model:   model,
			BaseURL: baseURLPtr,
		})
		if err != nil {
			return "", err
		}

		resp, genErr = cm.Generate(ctx, in)

	case string(commonModel.Gemini):
		client, err := genai.NewClient(ctx, &genai.ClientConfig{
			APIKey: apiKey,
		})
		if err != nil {
			return "", err
		}
		cm, err := gemini.NewChatModel(ctx, &gemini.Config{
			Client: client,
			Model:  model,
		})
		if err != nil {
			return "", err
		}

		resp, genErr = cm.Generate(ctx, in)

	case string(commonModel.Qwen):
		cm, err := qwen.NewChatModel(ctx, &qwen.ChatModelConfig{
			APIKey:  apiKey,
			Model:   setting.Model,
			BaseURL: baseURL,
		})
		if err != nil {
			return "", err
		}

		resp, genErr = cm.Generate(ctx, in)

	case string(commonModel.DeepSeek):
		cm, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
			APIKey:  apiKey,
			Model:   model,
			BaseURL: baseURL,
		})
		if err != nil {
			return "", err
		}

		resp, genErr = cm.Generate(ctx, in)

	case string(commonModel.Ollama):
		cm, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
			Model:   model,
			BaseURL: baseURL,
		})
		if err != nil {
			return "", err
		}

		resp, genErr = cm.Generate(ctx, in)

	case string(commonModel.Custom):
		cm, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
			APIKey:  apiKey,
			Model:   model,
			BaseURL: baseURL,
		})
		if err != nil {
			return "", err
		}

		resp, genErr = cm.Generate(ctx, in)

	default:
		return "", errors.New(commonModel.AGENT_PROVIDER_NOT_FOUND)
	}

	if genErr != nil {
		return "", genErr
	}

	return resp.Content, nil
}
