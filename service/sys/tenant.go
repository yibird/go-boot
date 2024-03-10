package sys

import (
	"github.com/gin-gonic/gin"
	"go-boot/global"
	modelCommon "go-boot/model/common"
	"go-boot/model/common/response"
	"go-boot/model/sys"
	"go-boot/service/common"
)

type TenantService struct {
	common.BaseService[sys.Tenant, sys.TenantDto]
}

var TenantServiceApp = new(TenantService)

func (s *TenantService) GetList(query sys.TenantQuery) (response.PageResult[sys.Tenant], error) {
	db := global.DB.Where("tenant_code like ?", "%"+query.TenantCode).
		Or("tenant_name like ?", "%"+query.TenantName).
		Where("super_tenant = ?", query.SuperTenant)
	page, err := response.SelectPage[sys.Tenant](db, query.Current, query.PageSize)
	return page, err
}

// Save 新增租户
func (s *TenantService) Save(c *gin.Context) error {
	tenant := &sys.Tenant{}
	tenantDto := &sys.TenantDto{}
	if err := modelCommon.Convert(c, tenantDto, tenant); err != nil {
		return nil
	}
	return TenantServiceApp.BaseService.Save(c, *tenant, *tenantDto)
}
