package routes

import (
	"it-mis/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Default(router *gin.Engine) {
	// 入口模块
	indexGroup := router.Group("/index")
	{
		index(indexGroup)
	}
	router.Use(middleware.LoginAuth())
	router.Use(middleware.CSRFTokenCheck())
	// 监控模块
	monitorGroup := router.Group("/apps/monitor")
	{
		monitor(monitorGroup)
	}
}
