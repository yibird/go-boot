package sys

import "go-boot/service"

type ApiGroup struct {
	AuthorityApi
	TenantApi
	RoleApi
}

var (
	tenantService = service.ServiceGroupApp.SystemServiceGroup.TenantService
	roleService   = service.ServiceGroupApp.SystemServiceGroup.RoleService
)
