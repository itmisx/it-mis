package db

import (
	"it-mis/init/config"
	"it-mis/internal/define"
	"it-mis/internal/model"
	"it-mis/internal/pkg/debugx"
	"it-mis/internal/pkg/encrypt"
	"it-mis/internal/pkg/randx"

	"gorm.io/gorm"
)

var _db *gorm.DB

type DB interface {
	InitDB(config.DBConfig) error
	NewDB() *gorm.DB
}

// InitDB 初始化数据库
func InitDB() {
	var db DB
	if config.Config.DBConfig.Type == "sqlite" {
		db = new(SQLite)
	} else if config.Config.DBConfig.Type == "mysql" {
		db = new(MySql)
	}
	if db.InitDB(config.Config.DBConfig) != nil {
		debugx.PrintPanic("init db error")
	}
	_db = db.NewDB()
}

// NewDB 新的数据库链接
func New() *gorm.DB {
	return _db
}

// AutoInstall 自动检测安装
func AutoInstall() {
	db := New()
	var count int64
	// 初始化version
	db.AutoMigrate(
		&model.Version{},
	)
	// 判断version表是否初始化
	db.Model(&model.Version{}).
		Count(&count)
	// verison表未初始化，则安装数据库
	if count <= 0 {
		// 安装数据库
		err := db.AutoMigrate(
			&model.User{},
			&model.UserToken{},
			&model.UserCode{},
			&model.UserSecureToken{},
			&model.SystemConfig{},
			&model.MonitorTask{},
		)
		if err != nil {
			debugx.PrintPanic(err)
		}
		// 初始化数据库版本
		db.Model(&model.Version{}).Create(&model.Version{
			ID:      1,
			Version: define.ModelVersion,
		})
		// 初始化管理员密码
		salt := randx.RandString(20)
		db.Model(&model.User{}).Create(&model.User{
			User:   "itmis",
			Salt:   salt,
			Passwd: encrypt.Sha256("Aa,123" + salt),
			Status: 1,
		})
	}
}
