package login

import (
	"it-mis/internal/pkg/randx"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SetSession 设置会话信息
func SetSession(c *gin.Context, key string, value interface{}) (err error) {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:     "/",
		Domain:   CookieDomain(c),
		MaxAge:   3600,
		HttpOnly: true,
	})
	session.Set(key, value)
	session.Save()
	return nil
}

// SetCSRFToken 设置csrf_token
func SetCSRFToken(c *gin.Context) (CSRFToken string) {
	// 设置csrf_token
	CSRFToken = randx.RandString(64)
	c.SetCookie("csrf_token", CSRFToken, 3600*24*30, "/", CookieDomain(c), false, true)
	return CSRFToken
}

func SetUserToken(c *gin.Context, userInfo UserInfo) error {
	// 设置csrf_token
	userToken := genUserToken(c, userInfo)
	if userToken != "" {
		c.SetCookie("user_token", userToken, 3600*24*30, "/", CookieDomain(c), false, true)
	}
	return nil
}

// CookieDomain 计算cookie作用域domain
func CookieDomain(c *gin.Context) (cookieDomain string) {
	domain := (strings.Split(c.Request.Host, ":"))[0]
	domainArr := strings.Split(domain, ".")
	if len(domainArr) > 1 {
		cookieDomain = "." + domainArr[len(domainArr)-2] + "." + domainArr[len(domainArr)-1]
	} else {
		cookieDomain = domainArr[0]
	}
	return cookieDomain
}
