package model

import (
	"it-mis/internal/pkg/uniqueid"

	"gorm.io/gorm"
)

// User 用户表
type User struct {
	ID     int64  `json:"id" gorm:"column:id;primaryKey;type:bigint(20) NOT NULL;comment:主键"`
	User   string `json:"user" gorm:"column:user;type:varchar(50) NOT NULL DEFAULT '';comment:用户名"`
	Passwd string `json:"passwd" gorm:"column:passwd;type:varchar(256) NOT NULL DEFAULT '';comment:密码"`
	Salt   string `json:"salt" gorm:"column:salt;type:varchar(256) NOT NULL DEFAULT '';comment:密码"`
	Status int    `json:"status" gorm:"default:1;column:status;type:tinyint(2) NOT NULL DEFAULT 1;commnet:状态"`
}

func (User) TableName() string {
	return "user"
}

// Create HOOK
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var UID int64
	UID, err = uniqueid.NewID()
	if err != nil {
		return err
	}
	u.ID = UID
	return
}
