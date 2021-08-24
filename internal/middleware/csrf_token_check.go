package middleware

import (
	"github.com/gin-gonic/gin"
)

// CSRFTokenCheck 跨站伪造检查
func CSRFTokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.GetHeader("CSRFToken")
		cookieToken, _ := c.Cookie("csrf_token")
		if headerToken != "" && cookieToken != "" && headerToken == cookieToken {
			c.Next()
		} else {
			c.Abort()
		}
	}
}
