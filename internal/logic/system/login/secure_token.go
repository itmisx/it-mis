package login

import (
	"it-mis/init/db"
	"it-mis/internal/model"
	"it-mis/internal/pkg/errorx"
	"time"

	"github.com/gin-gonic/gin"
)

// GetSecureToken 获取安全token
func (l Login) GetSecureToken(c *gin.Context) (secureToken string, err error) {
	var (
		userp = &model.User{}
		db    = db.New()
	)

	// 参数验证
	userp, err = l.validateParams(c, "secure_token")
	if err != nil {
		return "", err
	}
	// 生成secureToken
	token := ""
	count := db.Model(&model.UserSecureToken{}).
		Create(&model.UserSecureToken{
			UserID:     userp.ID,
			Token:      token,
			ExpireTime: time.Now().Unix() + 30*24*3600,
		}).
		RowsAffected
	if count == 1 {
		return token, nil
	}
	return "", errorx.New("生成安全令牌失败")
}

// GetSecureCode 获取安全码
func (l Login) GetSecureCode(c *gin.Context) (secureCode string, err error) {
	// 用户名、密码验证
	var (
		userInfo = &model.User{}
		db       = db.New()
	)

	// 参数验证
	userInfo, err = l.validateParams(c, "secure_token")
	if err != nil {
		return "", err
	}
	// 生成secureCode
	code := ""
	count := db.Model(&model.UserCode{}).
		Create(&model.UserCode{
			UserID:     userInfo.ID,
			Code:       code,
			ExpireTime: time.Now().Unix() + 180,
		}).
		RowsAffected
	if count == 1 {
		return code, nil
	}
	return "", errorx.New("生成安全码失败")
}
