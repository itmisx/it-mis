package model

import (
	"it-mis/internal/pkg/uniqueid"

	"gorm.io/gorm"
)

// MonitorTask 监控任务
type MonitorTask struct {
	ID          int64  `json:"id" gorm:"column:id;primaryKey;type:bigint(20) NOT NULL;comment:主键"`
	Name        string `json:"name" gorm:"column:name;type:varchar(50) NOT NULL DEFAULT '';comment:监控名称"`
	Host        string `json:"host" gorm:"column:host;type:varchar(255) NOT NULL DEFAULT '';comment:监控地址"`
	Port        string `json:"port" gorm:"column:port;type:int NOT NULL DEFAULT 0;comment:端口"`
	MonitorType int    `json:"monitor_type" gorm:"column:monitor_type;type:int(2) NOT NULL DEFAULT 1;comment:监控类型"`
	Interval    int    `json:"interval" gorm:"column:interval;type:int NOT NULL DEFAULT 1;comment:监控间隔"`
	Status      int    `json:"status" gorm:"column:status;type:int(2) NOT NULL DEFAULT 1;default 1;comment:状态，1-正常，2-异常"`
	Enable      int    `json:"enable" gorm:"column:enable;type:int(1) NOT NULL DEFAULT 1;comment:使能，1-启用，2-禁用"`
	// 多个异常信息以逗号分隔，前端用tag显示
	ExceptionInfo string `json:"exception_info" gorm:"column:exception_info;type:varchar(255) NOT NULL DEFAULT '';comment:异常信息"`
	CreateTime    int64  `json:"create_time" gorm:"column:create_time;type:bigint(20) NOT NULL DEFAULT 0;comment:创建时间"`
	UpdateTime    int64  `json:"update_time" gorm:"column:update_time;type:bigint(20) NOT NULL DEFAULT 0;comment:更新时间"`
}

func (MonitorTask) TableName() string {
	return "monitor_task"
}

// Create HOOK
func (u *MonitorTask) BeforeCreate(tx *gorm.DB) (err error) {
	var UID int64
	UID, err = uniqueid.NewID()
	if err != nil {
		return err
	}
	u.ID = UID
	return
}
