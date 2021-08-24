package routes

import "github.com/gin-gonic/gin"

func Default(router *gin.Engine) {
	// 入口模块
	indexGroup := router.Group("/index")
	{
		indexRoutes(indexGroup)
	}

}
