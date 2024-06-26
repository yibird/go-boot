package sys

import (
	"github.com/gin-gonic/gin"
	"go-boot/core/model/res"
)

type TenantApi struct {
	Service *TenantService
}

func NewTenantApi(Service *TenantService) *TenantApi {
	return &TenantApi{Service}
}

// GetList
// @Tags 	  Tenant
// @Summary   获取租户列表
// @accept    application/json
// @Produce   application/json
// @Success   200   {object} res.Response{data=object,msg=string}  "获取租户列表"
// @Router    /tenant/getList [get]
func (s *TenantApi) GetList(c *gin.Context) {
	var tenantQuery TenantQuery
	if err := c.ShouldBindQuery(&tenantQuery); err != nil {
		return
	}
	pageResult, err := s.Service.GetList(tenantQuery)
	res.ApiResultWithData(c, err, pageResult)
}

// Save 添加租户
func (s *TenantApi) Save(c *gin.Context) {
	res.SaveResult(c, s.Service.Save(c))
}

// Del 删除租户
func (s *TenantApi) Del(c *gin.Context) {

}

// Update 修改租户
func (s *TenantApi) Update(c *gin.Context) {

}
