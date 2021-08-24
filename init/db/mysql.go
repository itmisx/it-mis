package db

import (
	"it-mis/init/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _mysql_db *gorm.DB

type MySql struct {
}

func (m MySql) InitDB(cfg config.DBConfig) error {
	dsn := cfg.Username + ":" + cfg.Password +
		"@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Database + "?" +
		"charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		return err
	}
	if cfg.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)
	_mysql_db = db
	return nil
}
func (m MySql) NewDB() *gorm.DB {
	return _mysql_db
}
