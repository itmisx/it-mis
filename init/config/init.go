package config

import (
	"fmt"
	"mis/init/mysql"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var Config config

type config struct {
	App   App
	Mysql mysql.Config
}

// InitConfig 初始化配置
func InitConfig() {
	// 获取当前路径,然后定位配置文件的路径
	_, Dir, _, _ := runtime.Caller(0)
	pwd := filepath.Dir(Dir)
	confPath := path.Join(pwd + "/../../configs")
	// 解析app_config
	vp := viper.New()
	vp.SetConfigName("app")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(confPath)
	if err := vp.ReadInConfig(); err != nil {
		fmt.Println("read in config err:%", err)
	}
	vp.Unmarshal(&Config.App)
	// 解析mysql_config
	vp = viper.New()
	vp.SetConfigName("database")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(confPath)
	if err := vp.ReadInConfig(); err != nil {
		fmt.Println("read in config err:%", err)
	}
	vp.Unmarshal(&Config.Mysql)

	fmt.Printf("%v\n", Config)
}
