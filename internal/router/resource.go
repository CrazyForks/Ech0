package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/lin-snow/ech0/internal/di"
	_ "github.com/lin-snow/ech0/internal/swagger"
)

// setupResourceRoutes 设置资源路由
func setupResourceRoutes(appRouterGroup *AppRouterGroup, h *di.Handlers) {
	// Swagger UI
	appRouterGroup.ResourceGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	appRouterGroup.ResourceGroup.GET("/rss", h.CommonHandler.GetRss)
}
