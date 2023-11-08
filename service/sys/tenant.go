package sys

import (
	"go-boot/global"
	"go-boot/model/sys"
)

type TenantService struct{}

var TenantServiceApp = new(TenantService)

func (s *TenantService) GetAuthoritys() (list interface{}, err error) {
	var tanants []sys.Tanant
	err = global.DB.Find(&tanants).Error
	return tanants, err
}
