package sys

import (
	"github.com/gin-gonic/gin"
	auth "go-boot/module/sys/auth"
	menu "go-boot/module/sys/menu"
	post "go-boot/module/sys/post"
	role "go-boot/module/sys/role"
	tenant "go-boot/module/sys/tenant"
	user "go-boot/module/sys/user"
)

type SysRouter struct {
	AuthRouter   *auth.AuthRouter
	RoleRouter   *role.RoleRouter
	TenantRouter *tenant.TenantRouter
	PostRouter   *post.PostRouter
	MenuRouter   *menu.MenuRouter
	UserRouter   *user.UserRouter
}

func NewSysRouter(AuthRouter *auth.AuthRouter, RoleRouter *role.RoleRouter,
	TenantRouter *tenant.TenantRouter,
	PostRouter *post.PostRouter,
	MenuRouter *menu.MenuRouter,
	UserRouter *user.UserRouter,
) *SysRouter {
	return &SysRouter{AuthRouter, RoleRouter, TenantRouter,
		PostRouter, MenuRouter, UserRouter,
	}
}

// RegisterRoute 注册sys模块路由
func (s *SysRouter) RegisterRoute(router *gin.RouterGroup) {
	s.AuthRouter.RegisterRouter(router)
	s.RoleRouter.RegisterRouter(router)
	s.TenantRouter.RegisterRouter(router)
	s.PostRouter.RegisterRouter(router)
	s.MenuRouter.RegisterRouter(router)
	s.UserRouter.RegisterRouter(router)
}
