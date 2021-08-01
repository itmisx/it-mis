package controller

import (
	"it-mis/internal/app/host/logic/console"

	"github.com/gin-gonic/gin"
)

type Console struct {
}

// 终端连接
func (cs Console) Connect(c *gin.Context) {
	console.Console{}.Connect(c)
}
