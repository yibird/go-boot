package sys

import (
	"go-boot/global"
	"go-boot/model/common/response"
	"go-boot/model/sys"
)

type RoleService struct{}

var RoleServiceApp = new(RoleService)

func (s *RoleService) GetList(query sys.RoleQuery) (response.PageResult[sys.Role], error) {
	db := global.DB.Where("role_code like ? ", "%"+query.RoleCode).
		Or("role_name like ? ", "%"+query.RoleName)
	page, err := response.SelectPage[sys.Role](db, query.Current, query.PageSize)
	return page, err
}

func Save() {

}
