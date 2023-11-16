package query

import (
	"go-boot/global"
	"go-boot/types"
)

type TenantQuery struct {
	global.BaseQueryModel
	TenantCode  string         `form:"tenantCode" json:"tenantCode"`
	TenantName  string         `form:"tenantName" json:"tenantName"`
	SuperTenant *types.BitBool `form:"superTenant" json:"superTenant"`
}
