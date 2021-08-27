package start

import (
	"it-mis/init/config"
	"it-mis/init/db"
	"it-mis/init/session"
	"it-mis/internal/define/lang"
	"it-mis/internal/middleware"
	"it-mis/internal/pkg/i18n"
	"it-mis/internal/routes"

	"github.com/gin-gonic/gin"
)

func Run() error {
	// 初始化数据库
	db.InitDB()
	db.AutoInstall()

	// 加载多语言
	i18n.LoadLangPack(lang.LangCode)

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
