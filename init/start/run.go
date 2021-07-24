package start

import (
	"mis/internal/routes"

	"github.com/gin-gonic/gin"
)

func Run() error {
	r := gin.New()
	routes.Default(r)
	r.Run(":8080")
	return nil
}
