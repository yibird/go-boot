package initialize

import (
	docs "go-boot/docs"
	"go-boot/global"
	"go-boot/middleware"
	"go-boot/router"
	"go-boot/utils/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwaggerRouter(r *gin.Engine) {
	env := global.CONFIG.System.Env
	if env == "pro" {
		return
	}
	structs.Merge(docs.SwaggerInfo, global.CONFIG.Swagger)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

// 注册公开路由组
func RegisterPublicGroup(r *gin.Engine, routerPrefix string) {
	publicGroup := r.Group(routerPrefix)
	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		RegisterSwaggerRouter(r)
	}
}

// 注册私有路由组
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
	RegisterPublicGroup(r, routerPrefix)
	RegisterPrivateGroup(r, routerPrefix)
	global.LOGGER.Info("router register success🎉")
	return r
}
