package sys

import (
	"go-boot/global"
	"go-boot/utils/time"
)

type Tanant struct {
	global.BaseModel

	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	// 租金金额
	RentAmount int64 `gorm:"rent_amount" json:"rentAmount"`
	// 押金金额
	DepositAmount int64 `gorm:"deposit_amount" json:"depositAmount"`
	// 支付方式
	PaymentMethod int8 `gorm:"payment_method" json:"paymentMethod"`
	// 租约开始时间
	StartTime *time.LocalTime `gorm:"start_time" json:"startTime"`
	// 租约结束时间
	EndTime *time.LocalTime `gorm:"end_time" json:"endTime"`
	// 紧急联系人
	EmergencyContact string `gorm:"emergency_contact" json:"emergencyContact"`
}

func (Tanant) TableName() string {
	return "sys_tenant"
}
