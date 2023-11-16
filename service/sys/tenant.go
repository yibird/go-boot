package sys

import (
	"go-boot/global"
	"go-boot/model/common/response"
	"go-boot/model/sys/domain"
	"go-boot/model/sys/query"
)

type TenantService struct{}

var TenantServiceApp = new(TenantService)

func (s *TenantService) GetList(query query.TenantQuery) (response.PageResult[domain.TenantDo], error) {
	db := global.DB.Where("role_code like ? OR role_name like ? AND super_tenant= ?",
		"%"+query.TenantCode, query.TenantName, query.SuperTenant)
	page, err := response.SelectPage[domain.TenantDo](db, query.Current, query.PageSize)
	return page, err
}
