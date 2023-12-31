package core

import (
	"flag"
	"fmt"
	"go-boot/core/internal"
	"go-boot/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Viper 初始化Viper。配置优先级: 命令行 > 环境变量 > 默认值
func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	// 设置配置文件
	v.SetConfigFile(config)
	// 设置配置文件类型
	v.SetConfigType("yaml")
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 读取配置到全局对象
	v.Unmarshal(&global.CONFIG)
	// 监听配置文件变化
	v.WatchConfig()
	// 监听配置文件变化回调
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		//if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		//	fmt.Println(err)
		//}
	})
	return v
}
