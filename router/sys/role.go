package sys

import (
	"github.com/gin-gonic/gin"
	v1 "go-boot/api/v1"
)

type RoleRouter struct{}

func (s *TenantRouter) InitRoleRouter(Router *gin.RouterGroup) {
	router := Router.Group("role")
	roleApi := v1.ApiGroupApp.SysApiGroup.RoleApi
	{
		router.POST("save", roleApi.Save)
		router.GET("getList", roleApi.GetList)
	}
}
