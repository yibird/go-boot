package global

import (
	"go-boot/types"
	"gorm.io/gorm"
	"time"
)

// BaseModel 通用模型
type BaseModel struct {
	ID         int64            `gorm:"primary" json:"id" mapstructure:"id"`                              // 主键ID
	TenantId   int64            `gorm:"tenant_id" json:"tenantId" comment:"租户id" `                        // 租户id
	DataStatus *types.BitBool   `gorm:"data_status;default:0" json:"dataStatus" comment:"数据状态(0禁用,1启用)" ` // 数据状态
	Remark     string           `gorm:"remark" json:"remark"`                                             // 备注
	Version    int              `gorm:"version;default:0" json:"gorm"`                                    // 版本号
	Sort       int              `gorm:"sort" json:"sort"`                                                 // 排序序号
	Deleted    *types.BitBool   `gorm:"deleted;default:0" json:"deleted"`                                 // 删除标识 0未删除,1已删除
	Creator    string           `gorm:"creator" json:"creator"`                                           // 创建者
	CreateTime *types.LocalTime `gorm:"create_time" json:"createTime"`                                    // 创建时间
	Updater    string           `gorm:"updater" json:"updater"`                                           // 修改者
	UpdateTime *types.LocalTime `gorm:"update_time" json:"updateTime"`                                    // 更新时间
}

// BaseQueryModel 通用查询模型
type BaseQueryModel struct {
	Current  int64 `form:"current" json:"current"`   // 当前页数
	PageSize int64 `form:"pageSize" json:"pageSize"` // 每页显示条数
}

// BeforeCreate 插入前钩子
func (model *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	// 自动填充租户、创建时间、创建人等字段
	tx.Statement.SetColumn("tenant_id", 1)
	tx.Statement.SetColumn("creator", "admin")
	tx.Statement.SetColumn("create_time", types.LocalTime(time.Now()))
	return
}

// BeforeUpdate 更新前钩子
func BeforeUpdate(tx *gorm.DB) (err error) {
	// 自动更新时间、修改人等字段
	tx.Statement.SetColumn("updater", "admin")
	tx.Statement.SetColumn("update_time", types.LocalTime(time.Now()))
	return
}
