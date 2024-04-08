package core

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go-boot/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ConstructorParam struct {
	fx.In
	Viper  *viper.Viper
	Config *config.Config
	Logger *zap.Logger
	DB     *gorm.DB
	Redis  *redis.Client
}

type Core struct {
	Viper  *viper.Viper
	Config *config.Config
	Logger *zap.Logger
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewCore(p ConstructorParam) (*Core, error) {
	return &Core{
		Viper:  p.Viper,
		Config: p.Config,
		Logger: p.Logger,
		DB:     p.DB,
		Redis:  p.Redis,
	}, nil
}
