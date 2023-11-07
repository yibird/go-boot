package core

import (
	"fmt"
	"go-boot/global"
	"go-boot/initialize"
	"time"

	"go.uber.org/zap"
)

func RunServer() {
	r := initialize.Routers()
	addr := fmt.Sprintf(":%s", global.CONFIG.System.Addr)
	// 启动服务
	r.Run(addr)
	time.Sleep(10 * time.Microsecond)
	global.LOGGER.Info("go-boot started successfully, port:", zap.String("address", addr))
}
