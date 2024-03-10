package sys

import "go-boot/service"

type ApiGroup struct {
	AuthorityApi
	TenantApi
	RoleApi
	PostApi
	MenuApi
}

var (
	tenantService = service.ServiceGroupApp.SystemServiceGroup.TenantService
	roleService   = service.ServiceGroupApp.SystemServiceGroup.RoleService
	postService   = service.ServiceGroupApp.SystemServiceGroup.PostService
	menuService   = service.ServiceGroupApp.SystemServiceGroup.MenuService
)
