package apps

import (
	"it-mis/internal/logic/apps/monitor"
	"it-mis/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type MonitorController struct{}

// List 监控列表
func (MonitorController) TaskList(c *gin.Context) {
	count, list := monitor.Task{}.List(c)
	response.JSON(c, gin.H{
		"total": count,
		"list":  list,
	}, nil)
}

// Add 新增监控
func (MonitorController) AddTask(c *gin.Context) {
	response.JSON(c, nil, monitor.Task{}.Add(c))
}

// Edit 编辑监控
func (MonitorController) EditTask(c *gin.Context) {
}

// Delete 删除监控
func (MonitorController) DeleteTask(c *gin.Context) {
}
