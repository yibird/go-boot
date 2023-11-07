package initialize

import (
	"go-boot/global"
	"go-boot/middleware"
	"go-boot/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	env := global.CONFIG.System.Env

	// 如果是开发或测试环境则添加跨域中间件
	if env == "dev" || env == "test" {
		r.Use(middleware.Cors())
	}

	routerPrefix := global.CONFIG.System.RouterPrefix

	PublicGroup := r.Group(routerPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	systemRouter := router.RouterGroupApp.System
	PrivateGroup := r.Group(routerPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).
		Use(middleware.CasbinRbac())
	{
		systemRouter.InitAuthorityRouter(PrivateGroup)
	}
	global.LOGGER.Info("router register success")
	return r
}
