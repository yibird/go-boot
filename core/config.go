package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-boot/config"
)

func NewConfig(v *viper.Viper) *config.Config {
	// 设置配置文件
	v.SetConfigFile("config.yaml")
	// 设置配置文件类型
	v.SetConfigType("yaml")
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var config *config.Config
	// 读取配置到全局对象
	if err := v.Unmarshal(&config); err != nil {
		return nil
	}
	// 监听配置文件变化
	v.WatchConfig()
	// 监听配置文件变化回调
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
	})
	return config
}
