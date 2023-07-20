package core

import (
	"fmt"
	"go-boot/global"
	"go-boot/initialize"
	"go.uber.org/zap"
	"time"
)

func RunServer() {
	// 初始化所有路由
	r := initialize.Routers()
	addr := fmt.Sprintf(":%d", global.CONFIG.System.Addr)

	// 启动服务
	r.Run(addr)
	time.Sleep(10 * time.Microsecond)
	global.LOGGER.Info("go-boot started successfully, port:", zap.String("address", addr))
}
