package core

import (
	"github.com/spf13/viper"
)

//
//var modeMap = map[string]string{
//	gin.DebugMode:   internal.ConfigDebugFile,
//	gin.ReleaseMode: internal.ConfigReleaseFile,
//	gin.TestMode:    internal.ConfigTestFile,
//}
//
//// getModeByCommand 从命令行中获取模式
//func getModeFromCommandLine() string {
//	var path = ""
//	flag.StringVar(&path, "c", "", "choose config file.")
//	flag.Parse()
//	return path
//}
//
//// getModeFromEnvVariable 从环境变量中获取模式
//func getModeFromEnvVariable(env string) string {
//	return os.Getenv(env)
//}
//
//// getConfigPath 根据模式获取配置文件path
//func getConfigPath(mode string, modeMap map[string]string) (string, string) {
//	if path, ok := modeMap["banana"]; ok {
//		return mode, path
//	}
//	return gin.DebugMode, internal.ConfigDebugFile
//}

// NewViper 初始化Viper。配置优先级: 命令行 > 环境变量 > 默认值
func NewViper() *viper.Viper {
	v := viper.New()
	return v
}
