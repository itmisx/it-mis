package model

import (
	"it-mis/internal/pkg/uniqueid"

	"gorm.io/gorm"
)

// UserCode 用户令牌表
type UserCode struct {
	ID         int64  `json:"id" gorm:"column:id;primaryKey;type:bigint(20) NOT NULL;comment:主键"`
	Code       string `json:"code" gorm:"column:code;type:varchar(50) NOT NULL DEFAULT '';comment:令牌"`
	UserID     int64  `json:"user_id" gorm:"column:user_id;type:bigint(20) NOT NULL DEFAULT 0;comment:用户id"`
	ExpireTime int64  `json:"expire_time" gorm:"column:expire_time;type:bigint(20) NOT NULL DEFAULT 0;comment:过期时间"`
}

func (UserCode) TableName() string {
	return "user_code"
}

// Create HOOK
func (u *UserCode) BeforeCreate(tx *gorm.DB) (err error) {
	var UID int64
	UID, err = uniqueid.NewID()
	if err != nil {
		return err
	}
	u.ID = UID
	return
}
