package sys

import (
	"go-boot/global"
	"go-boot/types"
)

type Tenant struct {
	global.BaseModel
	TenantCode       string         `json:"tenantCode"`       // 租户编码
	TenantName       string         `json:"tenantName"`       // 租户名称
	SuperTenant      *types.BitBool `json:"superTenant"`      // 是否是超级租户(0否,1是)
	Address          string         `json:"address"`          // 地址
	Phone            string         `json:"phone"`            // 手机号
	Email            string         `json:"email"`            // 邮箱
	EmergencyContact string         `json:"emergencyContact"` // 紧急联系人
}

func (Tenant) TableName() string {
	return "sys_tenant"
}

type TenantDto struct {
}

type TenantQuery struct {
	global.BaseQueryModel
	TenantCode  string         `form:"tenantCode" json:"tenantCode"`
	TenantName  string         `form:"tenantName" json:"tenantName"`
	SuperTenant *types.BitBool `form:"superTenant" json:"superTenant"`
}
