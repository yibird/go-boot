package sys

import (
	"go-boot/core/model"
	"go-boot/core/model/req"
)

type Role struct {
	model.BaseModel
	// 租户id
	TenantId int64 `gorm:"tenant_id" json:"tenantId"`
	// 角色名称
	RoleName string `gorm:"role_name" json:"roleName"`
	// 角色权限标识
	RoleKey string `gorm:"role_key" json:"roleKey"`
	// 数据作用范围(0:全部数据权限,1:本部门数据权限,2:本部门及子机构数据权限,3:自定义数据权限)
	DataScope string `gorm:"data_scope" json:"dataScope"`
	// 状态(0启用,1禁用)
	Status int8 `gorm:"status" json:"status"`
}

func (Role) TableName() string {
	return "sys_role"
}

type RoleDto struct {
}
type RoleQuery struct {
	req.BaseQueryModel
	// 角色编码
	RoleCode string `gorm:"role_code" json:"roleCode"`
	// 角色名称
	RoleName string `gorm:"role_name" json:"roleName"`
}
