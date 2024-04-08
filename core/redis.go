package core

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-boot/config"
	"go.uber.org/zap"
)

// NewRedisClient 初始化redis.Client
func NewRedisClient(config *config.Config, logger *zap.Logger) *redis.Client {
	if config == nil {
		return nil
	}
	redisConfig := config.Redis
	// 创建redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Db,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("redis connect ping failed, err:", zap.Error(err))
		return nil
	}
	logger.Info("redis connect ping res:", zap.String("pong", pong))
	return client
}
