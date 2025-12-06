package event

import (
	"context"

	"github.com/lin-snow/ech0/internal/agent"
	echoRepository "github.com/lin-snow/ech0/internal/repository/echo"
	keyvalue "github.com/lin-snow/ech0/internal/repository/keyvalue"
	todoRepository "github.com/lin-snow/ech0/internal/repository/todo"
	userRepository "github.com/lin-snow/ech0/internal/repository/user"
)

type AgentProcessor struct {
	echoRepo     echoRepository.EchoRepositoryInterface
	todoRepo     todoRepository.TodoRepositoryInterface
	userRepo     userRepository.UserRepositoryInterface
	keyvalueRepo keyvalue.KeyValueRepositoryInterface
}

func NewAgentProcessor(
	echoRepo echoRepository.EchoRepositoryInterface,
	todoRepo todoRepository.TodoRepositoryInterface,
	userRepo userRepository.UserRepositoryInterface,
	keyvalueRepo keyvalue.KeyValueRepositoryInterface,
) *AgentProcessor {
	return &AgentProcessor{
		echoRepo:     echoRepo,
		todoRepo:     todoRepo,
		userRepo:     userRepo,
		keyvalueRepo: keyvalueRepo,
	}
}

func (ap *AgentProcessor) Handle(ctx context.Context, e *Event) error {
	// 清理生成内容的缓存
	// 删除 AGENT_GEN_RECENT 缓存
	ap.keyvalueRepo.DeleteKeyValue(context.Background(), string(agent.GEN_RECENT))

	return nil
}
