package sys

import (
	"go-boot/global"
	"go-boot/types"
)

type Menu struct {
	global.BaseModel
	ParentId      int            `json:"parent_id"`     // 菜单父id
	MenuCode      string         `json:"menuCode"`      // 菜单编码
	MenuName      string         `json:"menuName"`      // 菜单名称
	MenuType      int            `json:"menuType"`      // 菜单类型
	MenuPath      string         `json:"menuPath"`      // 菜单路径
	MenuIcon      string         `json:"menuIcon"`      // 菜单icon
	LevelPath     string         `json:"levelPath"`     // 菜单层级路径
	PermissionKey string         `json:"permissionKey"` // 权限标识
	ExternalLink  *types.BitBool `json:"externalLink"`  // 是否是外部链接
	Closable      *types.BitBool `json:"closable"`      // 是否可关闭
	IsNew         *types.BitBool `json:"isNew"`         // 是否是新菜单
}

func (Menu) TableName() string {
	return "sys_menu"
}

type MenuDto struct {
}

type MenuQuery struct {
	global.BaseQueryModel
	MenuCode string `form:"menuCode" json:"menuCode"`  // 菜单编码
	MenuName string ` form:"menuName" json:"menuName"` // 菜单名称
}
