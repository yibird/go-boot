package sys

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	TenantRouter
	AuthorityRouter
	PostRouter
}

func (r *RouterGroup) InitSysRouter(Router *gin.RouterGroup) {
	r.InitAuthorityRouter(Router)
	r.InitTenantRouter(Router)
	r.InitRoleRouter(Router)
	r.InitPostRouter(Router)
}
