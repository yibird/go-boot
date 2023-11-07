package main

import (
	"go-boot/core"
	"go-boot/global"
	"go-boot/initialize"
)

func main() {
	global.VIPER = core.Viper()
	global.LOGGER = core.Zap()
	global.DB = initialize.Gorm()
	core.RunServer()
}
