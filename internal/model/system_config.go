package model

import (
	"it-mis/internal/pkg/uniqueid"

	"gorm.io/gorm"
)

// SystemConfig 系统配置表
type SystemConfig struct {
	ID         int64  `json:"id" gorm:"column:id;primaryKey;type:bigint(20) NOT NULL;comment:主键"`
	Key        string `json:"key" gorm:"column:key;index:key;type:varchar(50) NOT NULL DEFAULT '';comment:配置项"`
	Value      string `json:"value" gorm:"column:type:varchar(1000) NOT NULL DEFAULT '';comment:配置内容"`
	EncrypType string `json:"encryp_type" gorm:"column:encry_type;type:varchar(30);comment:加密类型"` // 关键信息会加密，如采用rsa等
	CreateTime int64  `json:"create_time" gorm:"column:create_time;type:bigint(20) NOT NULL DEFAULT 0;comment:创建时间"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time;type:bigint(20) NOT NULL DEFAULT 0;comment:更新时间"`
}

func (SystemConfig) TableName() string {
	return "system_config"
}

// Create HOOK
func (s *SystemConfig) BeforeCreate(tx *gorm.DB) (err error) {
	var UID int64
	UID, err1 := uniqueid.NewID()
	if err1 != nil {
		return err
	}
	s.ID = UID
	return
}
