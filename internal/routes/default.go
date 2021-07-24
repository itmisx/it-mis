package routes

import "github.com/gin-gonic/gin"

func Default(router *gin.Engine) {
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
