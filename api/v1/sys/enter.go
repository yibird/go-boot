package sys

import "go-boot/service"

type ApiGroup struct {
	AuthorityApi
	TenantApi
}

var (
	tenantService = service.ServiceGroupApp.SystemServiceGroup.TenantService
)
