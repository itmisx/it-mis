package model

import (
	"it-mis/internal/pkg/uniqueid"

	"gorm.io/gorm"
)

// UserToken 用户安全令牌
type UserToken struct {
	ID         int64  `json:"id" gorm:"column:id;primaryKey;type:bigint(20) NOT NULL;comment:主键"`
	UserID     int64  `json:"user_id" gorm:"column:user_id;type:bigint(20);comment:用户id"`
	Token      string `json:"token" gorm:"column:token;type:varchar(50) NOT NULL DEFAULT '';comment:安全令牌"`
	ExpireTime int64  `json:"expire_time" gorm:"column:expire_time;type:bigint(20) NOT NULL DEFAULT 0;comment:过期时间"`
}

func (UserToken) TableName() string {
	return "user_token"
}

// Create HOOK
func (u *UserToken) BeforeCreate(tx *gorm.DB) (err error) {
	var UID int64
	UID, err = uniqueid.NewID()
	if err != nil {
		return err
	}
	u.ID = UID
	return
}
