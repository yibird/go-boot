package core

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-boot/core/internal"
	"os"
)

// NewViper 初始化Viper。配置优先级: 命令行 > 环境变量 > 默认值
func NewViper(path ...string) (*viper.Viper, string) {
	var configPath string
	if len(path) == 0 {
		flag.StringVar(&configPath, "c", "", "choose config file.")
		flag.Parse()
		if configPath == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				switch gin.Mode() {
				case gin.DebugMode:
					configPath = internal.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					configPath = internal.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					configPath = internal.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				configPath = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, configPath)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", configPath)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		configPath = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", configPath)
	}
	v := viper.New()
	return v, configPath
}
