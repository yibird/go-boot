package sys

import (
	"github.com/gin-gonic/gin"
)

type MenuRouter struct {
	Api *MenuApi
}

func NewMenuRouter(Api *MenuApi) *MenuRouter {
	return &MenuRouter{Api}
}

func (s *MenuRouter) RegisterRouter(Router *gin.RouterGroup) {
	router := Router.Group("menu")
	{
		router.POST("save", s.Api.Save)
		router.GET("getList", s.Api.GetList)
	}
}
