package sys

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	TenantRouter
	AuthorityRouter
}

func (r *RouterGroup) InitSysRouter(Router *gin.RouterGroup) {
	r.InitTenantRouter(Router)
	r.InitAuthorityRouter(Router)
}
