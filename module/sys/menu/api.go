package sys

import "github.com/gin-gonic/gin"

type MenuApi struct {
	Service *MenuService
}

func NewMenuApi(Service *MenuService) *MenuApi {
	return &MenuApi{Service}
}

// GetList 获取菜单列表
func (s *MenuApi) GetList(c *gin.Context) {

}

// Save 菜单
func (s *MenuApi) Save(c *gin.Context) {

}
