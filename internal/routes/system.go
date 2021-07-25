package routes

import "github.com/gin-gonic/gin"

func SystemRoutes(rg *gin.RouterGroup) {
	rg.POST("/sign-out")
}
