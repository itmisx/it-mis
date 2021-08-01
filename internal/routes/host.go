package routes

import (
	"it-mis/internal/app/host/controller"

	"github.com/gin-gonic/gin"
)

func hostRoutes(rg *gin.RouterGroup) {
	rg.Any("/console", controller.Console{}.Connect)
}
