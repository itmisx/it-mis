package db

import (
	"it-mis/init/config"
	"it-mis/internal/pkg/pathx"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _sqlite_db *gorm.DB

type SQLite struct{}

func (s SQLite) InitDB(cfg config.DBConfig) error {
	dbPath := path.Join(pathx.RootPath, "assets", "db", "itmis.db")
	// 初始化数据库存储目录
	{
		dbPath := path.Join(pathx.RootPath, "assets", "db")
		pathx.Mkdir(dbPath)
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if cfg.Debug {
		db = db.Debug()
	}
	if err != nil {
		return err
	}
	_sqlite_db = db
	return err
}

func (s SQLite) NewDB() *gorm.DB {
	return _sqlite_db
}
