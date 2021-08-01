package model

import (
	"it-mis/internal/pkg/uniqueid"

	"gorm.io/gorm"
)

// UserToken 用户令牌表
type UserToken struct {
	ID         int64  `json:"id" gorm:"column:id;primaryKey;type:bigint(20) NOT NULL;comment:主键"`
	Token      string `json:"token" gorm:"column:token;type:varchar(50) NOT NULL DEDAULT '';comment:令牌"`
	UserID     int64  `json:"user_id" gorm:"column:user_id;type:bigint(20) NOT NULL DEFAULT 0;comment:用户id"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time;type:bigint(20) NOT NULL DEFAULT 0;comment:创建时间"`
	ExpireTime int64  `json:"expire_time" gorm:"column:expire_time;type:bigint(20) NOT NULL DEFAULT 0;comment:过期时间"`
}

func (UserToken) TableName() string {
	return "user_token"
}

// Create HOOK
func (u *UserToken) BeforeCreate(tx *gorm.DB) (err error) {
	var UID int64
	UID, err1 := uniqueid.NewID()
	if err1 != nil {
		return err
	}
	u.ID = UID
	return
}
