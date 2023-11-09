package sys

import (
	res "go-boot/model/common/response"

	"github.com/gin-gonic/gin"
)

type TenantApi struct {
}

// 添加租户
func (s *TenantApi) AddTenant(c *gin.Context) {

}

// 删除租户
func (s *TenantApi) DelTenant(c *gin.Context) {

}

// 修改租户
func (s *TenantApi) UpTenant(c *gin.Context) {

}

// GetTenants
// @Tags 	  Tenant
// @Summary   获取租户列表
// @accept    application/json
// @Produce   application/json
// @Success   200   {object} response.Response{data=object,msg=string}  "获取租户列表"
// @Router    /tenant/GetTenants [get]
func (s *TenantApi) GetTenants(c *gin.Context) {
	list, err := tenantService.GetAuthoritys()
	res.ApiResultWithData(c, err, list)
}
