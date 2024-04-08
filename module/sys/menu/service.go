package sys

import (
	"go-boot/core/model/res"
	"go-boot/core/service"
	"gorm.io/gorm"
)

type MenuService struct {
	service.BaseService[Menu, MenuDto]
	DB *gorm.DB
}

func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{DB: db}
}

func (s *MenuService) GetList(query MenuQuery) (res.PageResult[Menu], error) {
	db := s.DB.Where("menu_code like ?", "%"+query.MenuCode).
		Or("menu_name like ?", "%"+query.MenuName)
	page, err := res.SelectPage[Menu](db, query.Current, query.PageSize)
	return page, err
}
