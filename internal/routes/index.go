package routes

import (
	"it-mis/internal/controller"
	"it-mis/internal/middleware"

	"github.com/gin-gonic/gin"
)

func indexRoutes(rg *gin.RouterGroup) {
	indexController := controller.IndexController{}
	rg.GET("/csrf-token", indexController.CSRFToken)
	rg.POST("/login", indexController.Login)
	rg.Use(middleware.LoginAuth())
	rg.Use(middleware.CSRFTokenCheck())
	rg.GET("/login-status", indexController.LoginStatus)
	rg.POST("/logout", indexController.Logout)
}
