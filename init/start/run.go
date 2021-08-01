package start

import (
	"it-mis/init/config"
	"it-mis/internal/pkg/debugx"
	"it-mis/internal/routes"

	"github.com/gin-gonic/gin"
)

func Run() error {
	// 自动生成rsa加密证书
	err := GenSecureRSA()
	if err != nil {
		debugx.PrintError(err)
	}
	// 启动gin
	r := gin.Default()
	routes.Default(r)
	r.Run(":" + config.Config.APP.BindPort)
	return nil
}
