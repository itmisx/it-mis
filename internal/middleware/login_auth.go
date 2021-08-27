package middleware

import (
	"it-mis/internal/logic/system/login"
	"it-mis/internal/pkg/errorx"
	"it-mis/internal/pkg/response"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginAuth 登录授权
func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查userToken
		cookie, _ := c.Cookie("user_token")
		if !login.ValidateUserToken(cookie) {
			response.JSON(c, nil,
				errorx.New("无效的user_token,请重新登录!", 200000, 1))
			c.Abort()
			return
		}
		// 检查session
		session := sessions.Default(c)
		userID := session.Get("user_id")
		if userID != nil {
			c.Next()
		} else {
			response.JSON(c, nil,
				errorx.New("无效的session,请重新登录！", 200000, 2))
			c.Abort()
			return
		}
	}
}
