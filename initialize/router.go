package initialize

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-boot/config"
	"go-boot/middleware"
	"go-boot/module/sys"
	"go-boot/utils/structs"
	"go.uber.org/zap"
	"net/http"
)

import (
	swaggerfiles "github.com/swaggo/files"
	docs "go-boot/docs"
)

func RegisterSwaggerRouter(r *gin.Engine, config *config.Config) {
	env := config.System.Env
	if env == "pro" {
		return
	}
	structs.Merge(docs.SwaggerInfo, config.Swagger)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

type EngineWrapper struct {
	Engine *gin.Engine
}

// NewEngine 创建*gin.Engine
func NewEngine() (*EngineWrapper, error) {
	r := gin.Default()
	return &EngineWrapper{r}, nil
}

func RegisterRoutes(ew *EngineWrapper, config *config.Config,
	logger *zap.Logger, sysRouter *sys.SysRouter) {
	engine := ew.Engine
	env := config.System.Env
	prefix := config.System.RouterPrefix
	// 如果是开发或测试环境则添加跨域中间件
	if env == "dev" || env == "test" {
		engine.Use(middleware.Cors())
	}

	publicGroup := engine.Group(config.System.RouterPrefix)
	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		RegisterSwaggerRouter(engine, config)
	}

	privateGroup := engine.Group(prefix)
	privateGroup.Use(
		middleware.Recover(),
		middleware.RateLimit(config.System.RateLimit),
		middleware.JWTAuth(),
		middleware.CasbinRbac(),
	)
	{
		sysRouter.RegisterRoute(privateGroup)
	}
	logger.Info("router register success🎉🎉🎉")
}
