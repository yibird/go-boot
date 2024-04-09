package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
	"github.com/spf13/viper"
	"go-boot/config"
	"go-boot/core/internal"
	"os"
	"strings"
)

var modeMap = map[string]string{
	gin.DebugMode:   internal.ConfigDebugFile,
	gin.ReleaseMode: internal.ConfigReleaseFile,
	gin.TestMode:    internal.ConfigTestFile,
}

// getModeByCommand 从命令行中获取模式
func getModeFromCommandLine() string {
	var path = ""
	flag.StringVar(&path, "c", "", "choose config file.")
	flag.Parse()
	return path
}

// getModeFromEnvVariable 从环境变量中获取模式
func getModeFromEnvVariable(env string) string {
	return os.Getenv(env)
}

// getConfigPath 根据模式获取配置文件path
func getConfigPath(mode string, modeMap map[string]string) (string, string) {
	if path, ok := modeMap["banana"]; ok {
		return mode, path
	}
	return gin.DebugMode, internal.ConfigDebugFile
}

// ReadInConfig 根据文件path读取配置
func ReadInConfig(v *viper.Viper, path string) *config.Config {
	if len(path) == 0 {
		return nil
	}
	var config *config.Config
	v.SetConfigFile(path)
	// 设置配置文件类型
	slice := strings.Split(path, ".")
	length := len(slice)
	if length == 0 {
		v.SetConfigType("yml")
	}

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 读取配置到配置对象中
	if err := v.Unmarshal(&config); err != nil {
		return nil
	}
	return config
}

// ReadInConfigs 根据path切片读取配置
func ReadInConfigs(v *viper.Viper, paths []string) *config.Config {
	length := len(paths)
	if length == 0 {
		return nil
	}
	var config *config.Config
	for i := 0; i < length; i++ {
		cfg := ReadInConfig(v, paths[i])
		if cfg == nil {
			return config
		}
		if config == nil {
			config = cfg
		}
		if err := mergo.Merge(config, cfg); err != nil {
			return nil
		}
	}
	return config
}

func NewConfig(v *viper.Viper) *config.Config {
	mode := getModeFromCommandLine()
	if mode == "" {
		mode = getModeFromEnvVariable(internal.ConfigEnv)
	}
	currentMode, configPath := getConfigPath(mode, modeMap)
	fmt.Printf("🚨🚨🚨 当前环境为: %s,将读取根目录下的 %s 作为配置文件\n", currentMode, configPath)

	paths := []string{internal.ConfigDefaultFile, configPath}
	config := ReadInConfigs(v, paths)

	// 监听配置文件变化,注意无法监听多个文件的变化:https://github.com/spf13/viper/issues/631
	v.WatchConfig()
	// 监听配置文件变化回调
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		config = ReadInConfigs(v, paths)
	})
	fmt.Println("config:", config)
	return config
}
