package handler

import service "github.com/lin-snow/ech0/internal/service/agent"

type AgentHandler struct {
	agentService service.AgentServiceInterface
}

func NewAgentHandler(agentService service.AgentServiceInterface) *AgentHandler {
	return &AgentHandler{
		agentService: agentService,
	}
}
