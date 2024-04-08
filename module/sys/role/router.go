package sys

import (
	"github.com/gin-gonic/gin"
)

type RoleRouter struct {
	RoleApi *RoleApi
}

func NewRoleRouter(RoleApi *RoleApi) *RoleRouter {
	return &RoleRouter{RoleApi}
}

func (s *RoleRouter) RegisterRouter(Router *gin.RouterGroup) {
	router := Router.Group("role")
	roleApi := s.RoleApi
	{
		router.POST("save", roleApi.Save)
		router.GET("getList", roleApi.GetList)
	}
}
