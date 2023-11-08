package sys

import (
	v1 "go-boot/api/v1"

	"github.com/gin-gonic/gin"
)

type TenantRouter struct{}

func (s *TenantRouter) InitTenantRouter(Router *gin.RouterGroup) {
	router := Router.Group("tenant")
	tenantApi := v1.ApiGroupApp.SysApiGroup.TenantApi
	{
		router.GET("getTenants", tenantApi.GetTenants)
	}
}
