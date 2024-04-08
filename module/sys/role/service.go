package sys

import (
	"go-boot/core/model/res"
	"gorm.io/gorm"
)

type RoleService struct {
	DB *gorm.DB
}

func NewRoleService(DB *gorm.DB) *RoleService {
	return &RoleService{DB}
}

// GetList 获取角色列表
func (s *RoleService) GetList(query RoleQuery) (res.PageResult[Role], error) {
	db := s.DB.Where("role_code like ? ", "%"+query.RoleCode).
		Or("role_name like ? ", "%"+query.RoleName)
	page, err := res.SelectPage[Role](db, query.Current, query.PageSize)
	return page, err
}

func Save() {

}
