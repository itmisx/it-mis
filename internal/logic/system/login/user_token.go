package login

import (
	"it-mis/init/db"
	"it-mis/internal/model"
	"it-mis/internal/pkg/randx"
	"time"

	"github.com/gin-gonic/gin"
)

// genUserToken 生成userToken
func genUserToken(c *gin.Context, userID int64) string {
	token := randx.RandString(128)
	count := db.New().Model(&model.UserToken{}).
		Create(&model.UserToken{
			UserID:     userID,
			Token:      token,
			ExpireTime: time.Now().Unix() + 3600*24*30,
		}).RowsAffected
	if count > 0 {
		return token
	}
	return ""
}

// ValidateUserToken 验证userToken的有效性
func ValidateUserToken(token string) bool {
	var count int64
	db.New().Model(&model.UserToken{}).
		Where("token = ?", token).
		Count(&count)
	return count > 0
}
