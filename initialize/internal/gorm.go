package internal

import (
	"go-boot/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type _gorm struct{}

var Gorm = new(_gorm)

// DBBase 数据库通用接口
type DBBase interface {
	// GetLoggerMode 获取日志模式
	GetLoggerMode() string
}

// Config 自定义GORM配置
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	// Gorm配置
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	// Gorm日志配置
	_logger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	var loggerMode DBBase

	switch global.CONFIG.System.DbType {
	case "mysql":
		loggerMode = &global.CONFIG.Mysql
	//case "pgsql":
	//	loggerMode = &global.CONFIG.Pgsql
	//case "oracle":
	//	loggerMode = &global.CONFIG.Oracle
	default:
		loggerMode = &global.CONFIG.Mysql
	}

	switch loggerMode.GetLoggerMode() {
	case "silent", "Silent":
		config.Logger = _logger.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _logger.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _logger.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _logger.LogMode(logger.Info)
	default:
		config.Logger = _logger.LogMode(logger.Info)
	}
	return config
}
