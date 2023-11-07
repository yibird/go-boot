package sys

import (
	v1 "go-boot/api/v1"

	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	router := Router.Group("authority")
	authorityApi := v1.ApiGroupApp.SysApiGroup.AuthorityApi
	{
		router.POST("createAuthority", authorityApi.CreateAuthority)
		router.GET("getAuthoritys", authorityApi.GetAuthoritys)
	}
}
