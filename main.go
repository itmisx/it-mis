package main

import (
	"mis/cmd"
	"mis/init/config"
)

func main() {
	// 初始化配置项
	config.InitConfig()
	// 命令执行
	cmd.Execute()
}
