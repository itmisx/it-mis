package start

import (
	"it-mis/init/config"
	"it-mis/init/db"
	"it-mis/init/session"
	"it-mis/internal/middleware"
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
	// 初始化数据库
	db.InitDB()
	db.AutoInstall()

	// 启动gin
	r := gin.Default()

	// 加载session
	sess := session.InitSession()
	r.Use(sess)

	// 加载中间件
	r.Use(middleware.SetCors())

	routes.Default(r)

	r.Run(":" + config.Config.APP.BindPort)
	return nil
}
