package global

import (
	"go-boot/utils/time"
)

// 通用模型
type BaseModel struct {
	ID         uint            `gorm:"primary" json:"id"`             // 主键ID
	Remark     string          `json:"remark"`                        // 备注
	Version    int             `json:"version"`                       // 版本号
	OrderNum   int             `gorm:"order_num" json:"orderNum"`     // 排序序号
	Deleted    int8            `gorm:"deleted" json:"deleted"`        // 删除标识 0未删除,1已删除
	Creator    string          `gorm:"creator" json:"creator"`        // 创建者
	CreateTime *time.LocalTime `gorm:"create_time" json:"createTime"` // 创建时间
	Updater    string          `gorm:"updater" json:"updater"`        // 修改者
	UpdateTime *time.LocalTime `gorm:"update_time" json:"updateTime"` // 更新时间
}
