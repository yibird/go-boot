package tenant

import (
	"go-boot/global"
	"time"
)

type TanantModel struct {
	global.BaseModel

	Name    string
	Address string
	Phone   string
	Email   string
	// 租金金额
	RentAmount int64 `gorm:"rent_amount" json:"rentAmount"`
	// 押金金额
	DepositAmount int64 `gorm:"deposit_amount" json:"depositAmount"`
	// 支付方式
	PaymentMethod int8 `gorm:"payment_method" json:"paymentMethod"`
	// 租约开始时间
	StartTime time.Time `gorm:"start_time" json:"startTime"`
	// 租约结束时间
	EndTime time.Time `gorm:"end_time" json:"endTime"`
	// 紧急联系人
	EmergencyContact string `gorm:"emergency_contact" json:"emergencyContact"`
}
