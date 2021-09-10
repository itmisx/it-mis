package routes

import (
	"it-mis/internal/controller/apps"

	"github.com/gin-gonic/gin"
)

func monitor(rg *gin.RouterGroup) {
	monitorController := apps.MonitorController{}
	rg.GET("/task-list", monitorController.TaskList)
	rg.POST("/add-task", monitorController.AddTask)
}
