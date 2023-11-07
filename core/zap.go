package core

import (
	"fmt"
	"go-boot/core/internal"
	"go-boot/global"

	"github.com/duke-git/lancet/v2/fileutil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 初始化Zap,获取zap.Logger实例
func Zap() (logger *zap.Logger) {
	director := global.CONFIG.Zap.Director
	if exist := fileutil.IsExist(director); !exist {
		fmt.Printf("create %v directory\n", director)
		fileutil.CreateDir(director)
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
