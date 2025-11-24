package service

type AgentService struct {
	// 实现 Agent 服务的方法
}

func NewAgentService() AgentServiceInterface {
	return &AgentService{}
}
