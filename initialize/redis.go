package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-boot/global"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.CONFIG.Redis
	// 创建redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Db,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOGGER.Error("redis connect ping failed, err:", zap.Error(err))
		return
	}
	global.LOGGER.Info("redis connect ping response:", zap.String("pong", pong))
	global.REDIS = client
}
