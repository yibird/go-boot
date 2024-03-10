package sys

import (
	"github.com/gin-gonic/gin"
	v1 "go-boot/api/v1"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	router := Router.Group("tenant")
	api := v1.ApiGroupApp.SysApiGroup.MenuApi
	{
		router.POST("save", api.Save)
		router.GET("getList", api.GetList)
	}
}
