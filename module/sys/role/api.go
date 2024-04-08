package sys

import (
	"github.com/gin-gonic/gin"
	"go-boot/core/model/res"
)

type RoleApi struct {
	Service *RoleService
}

func NewRoleApi(Service *RoleService) *RoleApi {
	return &RoleApi{Service}
}

func (s *RoleApi) Save(c *gin.Context) {

}

// GetList
// @Tags 	  Role
// @Summary   获取角色列表
// @accept    application/json
// @Produce   application/json
// @Success   200   {object} res.Response{data=object,msg=string}  "获取角色列表"
// @Router    /role/getList [get]
func (s *RoleApi) GetList(c *gin.Context) {
	var roleQuery RoleQuery
	if err := c.ShouldBindQuery(&roleQuery); err != nil {
		return
	}
	pageResult, err := s.Service.GetList(roleQuery)
	res.ApiResultWithData(c, err, pageResult)
}
