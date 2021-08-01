package routes

import "github.com/gin-gonic/gin"

func Default(router *gin.Engine) {
	// 入口模块
	indexGroup := router.Group("/index")
	{
		indexRoutes(indexGroup)
	}
	//
	systemGroup := router.Group("/auth")
	{
		SystemRoutes(systemGroup)
	}
	// 项目管理
	// projectGroup := router.Group("/project")
	// {
	// 	projectRoutes(projectGroup)
	// }
	// 主机管理
	hostGroup := router.Group("/host")
	{
		hostRoutes(hostGroup)
	}

}
