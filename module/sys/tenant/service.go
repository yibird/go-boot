package sys

import (
	"github.com/gin-gonic/gin"
	"go-boot/core/convert"
	"go-boot/core/model/res"
	"go-boot/core/service"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TenantService struct {
	service.BaseService[Tenant, TenantDto]
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewTenantService(db *gorm.DB, logger *zap.Logger) *TenantService {
	return &TenantService{DB: db, Logger: logger}
}

var TenantServiceApp = new(TenantService)

func (s *TenantService) GetList(query TenantQuery) (res.PageResult[Tenant], error) {
	db := s.DB.Where("tenant_code like ?", "%"+query.TenantCode).
		Or("tenant_name like ?", "%"+query.TenantName).
		Where("super_tenant = ?", query.SuperTenant)
	page, err := res.SelectPage[Tenant](db, query.Current, query.PageSize)
	return page, err
}

// Save 新增租户
func (s *TenantService) Save(c *gin.Context) error {
	tenant := &Tenant{}
	tenantDto := &TenantDto{}
	if err := convert.Convert(c, tenantDto, tenant); err != nil {
		return nil
	}
	return TenantServiceApp.BaseService.Save(c, *tenant, *tenantDto)
}
