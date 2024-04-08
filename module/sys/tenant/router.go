package sys

import (
	"github.com/gin-gonic/gin"
)

type TenantRouter struct {
	Api *TenantApi
}

func NewTenantRouter(Api *TenantApi) *TenantRouter {
	return &TenantRouter{Api}
}

func (s *TenantRouter) RegisterRouter(Router *gin.RouterGroup) {
	router := Router.Group("tenant")
	api := s.Api
	{
		router.POST("save", api.Save)
		router.GET("getList", api.GetList)
	}
}
