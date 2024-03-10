package sys

import (
	"github.com/gin-gonic/gin"
	v1 "go-boot/api/v1"
)

type PostRouter struct{}

func (s *PostRouter) InitPostRouter(Router *gin.RouterGroup) {
	router := Router.Group("post")
	api := v1.ApiGroupApp.SysApiGroup.PostApi
	{
		router.GET("getList", api.GetList)
		router.POST("save", api.Save)
		router.POST("delByIds", api.DelByIds)
		router.POST("update", api.Update)
	}
}
