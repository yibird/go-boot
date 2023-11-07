package global

import (
	"time"
)

// 通用模型
type BaseModel struct {
	ID         uint      `gorm:"primary"` // 主键ID
	Remark     string    // 备注
	version    int       // 版本号
	OrderNum   int       `gorm:"order_num"` // 排序序号
	Deleted    int8      // 删除标识 0未删除,1已删除
	Creator    string    // 创建者
	CreateTime time.Time `gorm:"create_time"` // 创建时间
	Updater    string    // 修改者
	UpdateTime time.Time `gorm:"update_time"` // 更新时间
}
