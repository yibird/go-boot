package domain

import "go-boot/global"

type RoleDo struct {
	global.BaseModel
	RoleCode  string `gorm:"role_code" json:"roleCode"`
	RoleName  string `gorm:"role_name" json:"roleName"`
	RoleKey   string `gorm:"role_key" json:"roleKey"`
	DataScope int    `gorm:"data_scope" json:"dataScope"`
}

func (RoleDo) TableName() string {
	return "sys_role"
}
