package login

import (
	"it-mis/init/db"
	"it-mis/internal/model"
	"it-mis/internal/pkg/randx"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetCSRFToken 设置csrf_token
func SetCSRFToken(c *gin.Context) (CSRFToken string) {
	// 设置csrf_token
	CSRFToken = randx.RandString(64)
	c.SetCookie("csrf_token", CSRFToken, 0, "/", CookieDomain(c), false, true)
	return CSRFToken
}

func SetUserToken(c *gin.Context, userID int64) (string, error) {
	// 设置csrf_token
	userToken := genUserToken(c, userID)
	if userToken != "" {
		c.SetCookie("user_token", userToken, 0, "/", CookieDomain(c), false, true)
	}
	return userToken, nil
}

func GetUserTokenInfo(userToken string) (tokenInfo *model.UserToken) {
	db := db.New()
	tokenInfo = &model.UserToken{}
	db.Model(&model.UserToken{}).
		Where("token = ?", userToken).
		Take(&tokenInfo)
	if tokenInfo.ID <= 0 {
		return nil
	}
	return tokenInfo
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
