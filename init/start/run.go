package start

import (
	"mis/internal/pkg"
	"mis/internal/routes"

	"github.com/gin-gonic/gin"
)

func Run() error {
	// 自动生成rsa加密证书
	err := GenSecureRSA()
	if err != nil {
		pkg.PrintError(err)
	}
	// 启动gin
	r := gin.Default()
	routes.Default(r)
	r.Run(":8080")
	return nil
}
