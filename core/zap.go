package core

import (
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"go-boot/config"
	"go-boot/core/internal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger 初始化Zap.Logger
func NewLogger(config *config.Config) (logger *zap.Logger) {
	// 获取日志目录
	director := config.Zap.Director
	if exist := fileutil.IsExist(director); !exist {
		fmt.Printf("create %v directory\n", director)
		if err := fileutil.CreateDir(director); err != nil {
			return nil
		}
	}
	cores := internal.Zap.GetZapCores(config)
	logger = zap.New(zapcore.NewTee(cores...))

	if config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
