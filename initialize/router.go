package initialize

import (
	"go-boot/global"
	"go-boot/middleware"
	"go-boot/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPublicGroup(r *gin.Engine, routerPrefix string) {
	publicGroup := r.Group(routerPrefix)
	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
}

func RegisterPrivateGroup(r *gin.Engine, routerPrefix string) {
	privateGroup := r.Group(routerPrefix)

	sysRouter := router.RouterGroupApp.System

	privateGroup.Use(middleware.RateLimit(global.CONFIG.System.RateLimit)).
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinRbac())
	{
		sysRouter.InitSysRouter(privateGroup)
	}

}

func Routers() *gin.Engine {
	r := gin.Default()
	env := global.CONFIG.System.Env

	// 如果是开发或测试环境则添加跨域中间件
	if env == "dev" || env == "test" {
		r.Use(middleware.Cors())
	}
	// 获取路由前缀
	routerPrefix := global.CONFIG.System.RouterPrefix
	// 注册公共路由组
	RegisterPublicGroup(r, routerPrefix)
	// 注册私有路由组
	RegisterPrivateGroup(r, routerPrefix)
	global.LOGGER.Info("router register success")
	return r
}
