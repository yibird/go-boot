package main

import (
	"go-boot/core"
	"go-boot/global"
	"go-boot/initialize"
)

func main() {
	// 初始化Viper
	global.VIPER = core.Viper()
	// 初始化Zap
	global.LOGGER = core.Zap()
	// 初始化Gorm
	global.DB = initialize.Gorm()
	// 运行服务器
	core.RunServer()
}
