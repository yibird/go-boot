package sys

import (
	"go-boot/global"
	"go-boot/model/common/response"
	"go-boot/model/sys"
	"go-boot/service/common"
)

type MenuService struct {
	common.BaseService[sys.Menu, sys.MenuDto]
}

var MenuServiceApp = new(MenuService)

func (s *MenuService) GetList(query sys.MenuQuery) (response.PageResult[sys.Menu], error) {
	db := global.DB.Where("menu_code like ?", "%"+query.MenuCode).
		Or("menu_name like ?", "%"+query.MenuName)
	page, err := response.SelectPage[sys.Menu](db, query.Current, query.PageSize)
	return page, err
}
