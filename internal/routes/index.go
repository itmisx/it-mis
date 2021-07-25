package routes

import (
	"mis/internal/app/system"

	"github.com/gin-gonic/gin"
)

func indexRoutes(rg *gin.RouterGroup) {
	rg.GET("/sign-in", system.AccessAuth{}.SignIn)
}
