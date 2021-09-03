package monitor

import "github.com/gin-gonic/gin"

type Task struct{}

// 任务列表
func (t Task) List(c *gin.Context) {
}

// 新建任务
func (t Task) Add(c *gin.Context) {
}

// 编辑任务
func (t Task) Edit(c *gin.Context) {
}

// 使能任务
func (t Task) Enable(c *gin.Context) {
}

// 删除任务
func (t Task) Delete(c *gin.Context) {
}
