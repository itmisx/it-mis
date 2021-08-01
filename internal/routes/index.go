package routes

import (
	"it-mis/internal/app/system/controller"

	"github.com/gin-gonic/gin"
)

func indexRoutes(rg *gin.RouterGroup) {
	login := controller.Login{}
	rg.GET("/login", login.Login)
	rg.GET("/logout", login.Logout)
	rg.GET("/status", login.Status)
}
