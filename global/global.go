package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go-boot/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	LOGGER *zap.Logger
	VIPER  *viper.Viper
	CONFIG config.Config
	REDIS  *redis.Client
)
