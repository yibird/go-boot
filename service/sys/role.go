package sys

import (
	"go-boot/global"
	"go-boot/model/common/response"
	"go-boot/model/sys/domain"
	"go-boot/model/sys/query"
)

type RoleService struct{}

var RoleServiceApp = new(RoleService)

func (s *RoleService) GetList(query query.RoleQuery) (response.PageResult[domain.RoleDo], error) {
	db := global.DB
	if query.RoleCode != "" {
		db = db.Where("role_code like ? ", query.RoleCode+"%")
	}
	if query.RoleName != "" {
		db = db.Or("role_name like ? ", query.RoleName+"%")
	}
	page, err := response.SelectPage[domain.RoleDo](db, query.Current, query.PageSize)
	return page, err
}
