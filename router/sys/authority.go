package sys

import (
	"github.com/gin-gonic/gin"
	v1 "go-boot/api/v1"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	router := Router.Group("authority")
	authorityApi := v1.ApiGroupApp.SysApiGroup.AuthorityApi
	{
		router.POST("createAuthority", authorityApi.CreateAuthority)
	}
}
