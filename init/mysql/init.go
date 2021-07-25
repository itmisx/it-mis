package mysql

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	_db *gorm.DB
)

type Config struct {
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	Host           string `mapstructure:"host" default:"127.0.0.1"`
	Port           string `mapstructure:"port" default:"3306"`
	Database       string `mapstructure:"database"`
	ConnectTimeout int    `mapstructure:"connect_timeout" default:"10"`
	Debug          bool   `mapstructure:"debug" default:"true"`
	MaxLifetime    int    `mapstructure:"max_lifetime" default:"7200"`  //设置连接可以重用的最长时间(单位：秒)
	MaxOpenConns   int    `mapstructure:"max_open_conns" default:"150"` //设置数据库的最大打开连接数
	MaxIdleConns   int    `mapstructure:"max_idle_conns" default:"50"`  //设置空闲连接池中的最大连接数
}

func Open(cfg Config) (*gorm.DB, error) {
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
		return nil, err
	}
	if cfg.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)
	return db, nil
}

// InitMysql mysql数据库初始化
func InitMysql(cfg Config) error {
	db, err := Open(cfg)
	if err != nil {
		return err
	}
	_db = db
	return nil
}

// 获取一个数据库对象
func NewDB() *gorm.DB {
	return _db
}
