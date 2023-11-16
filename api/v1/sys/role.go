package sys

import (
	"github.com/gin-gonic/gin"
	res "go-boot/model/common/response"
	"go-boot/model/sys/query"
)

type RoleApi struct {
}

func (s *RoleApi) Save(c *gin.Context) {

}

// GetList
// @Tags 	  Role
// @Summary   获取角色列表
// @accept    application/json
// @Produce   application/json
// @Success   200   {object} response.Response{data=object,msg=string}  "获取角色列表"
// @Router    /role/getList [get]
func (s *RoleApi) GetList(c *gin.Context) {
	var roleQuery query.RoleQuery
	if err := c.ShouldBindQuery(&roleQuery); err != nil {
		return
	}
	pageResult, err := roleService.GetList(roleQuery)
	res.ApiResultWithData(c, err, pageResult)
}
