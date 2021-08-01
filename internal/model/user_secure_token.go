package model

import (
	"it-mis/internal/pkg/uniqueid"

	"gorm.io/gorm"
)

// UserSecureToken 用户安全令牌
type UserSecureToken struct {
	ID         int64  `json:"id" gorm:"column:id;primaryKey;type:bigint(20) NOT NULL;comment:主键"`
	UserID     int64  `json:"user_id" gorm:"column:user_id;type:bigint(20);comment:用户id"`
	Token      string `json:"token" gorm:"column:token;type:varchar(50) NOT NULL DEFAULT '';comment:安全令牌"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time;type:bigint(20) NOT NULL DEFAULT 0;comment:创建时间"`
}

func (UserSecureToken) TableName() string {
	return "user_secure_token"
}

// Create HOOK
func (u *UserSecureToken) BeforeCreate(tx *gorm.DB) (err error) {
	var UID int64
	UID, err1 := uniqueid.NewID()
	if err1 != nil {
		return err
	}
	u.ID = UID
	return
}
