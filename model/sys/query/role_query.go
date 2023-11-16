package query

import "go-boot/global"

type RoleQuery struct {
	global.BaseQueryModel
	RoleCode string `form:"roleCode" json:"roleCode"`
	RoleName string `form:"roleName" json:"roleName"`
}
