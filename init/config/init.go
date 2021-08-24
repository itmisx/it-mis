package config

import (
	"fmt"
	"it-mis/internal/pkg/debugx"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var Config config

type config struct {
	RootPath string `json:"root_path" mapstructure:"root_path"`
	APP      APP
	DBConfig DBConfig
	Email    Email
}

// InitConfig 初始化配置
func InitConfig() {
	// 获取当前路径,然后定位配置文件的路径
	// 以适用单元测试
	_, Dir, _, _ := runtime.Caller(0)
	pwd := filepath.Dir(Dir)
	Config.RootPath = path.Join(pwd + "/../../")
	confPath := path.Join(Config.RootPath + "/configs")

	// 解析app_config
	vp := viper.New()
	vp.SetConfigName("app")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(confPath)
	if err := vp.ReadInConfig(); err != nil {
		fmt.Println("read in config err:%", err)
	}
	vp.Unmarshal(&Config.APP)
	Config.APP.RSAPath = path.Join(Config.RootPath, Config.APP.RSAPath)

	// 解析mysql_config
	vp = viper.New()
	vp.SetConfigName("database")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(confPath)
	if err := vp.ReadInConfig(); err != nil {
		fmt.Println("read in config err:%", err)
	}
	vp.Unmarshal(&Config.DBConfig)
	debugx.PrintInfo(Config)
}
