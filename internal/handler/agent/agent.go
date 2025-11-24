package handler

import (
	"github.com/gin-gonic/gin"

	res "github.com/lin-snow/ech0/internal/handler/response"
	service "github.com/lin-snow/ech0/internal/service/agent"
)

type AgentHandler struct {
	agentService service.AgentServiceInterface
}

func NewAgentHandler(agentService service.AgentServiceInterface) *AgentHandler {
	return &AgentHandler{
		agentService: agentService,
	}
}

func (agentHandler *AgentHandler) GetRecent() gin.HandlerFunc {
	return res.Execute(func(ctx *gin.Context) res.Response {
		return res.Response{
			Msg: "获取作者近况成功",
		}
	})
}
