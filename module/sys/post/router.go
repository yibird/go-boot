package sys

import (
	"github.com/gin-gonic/gin"
)

type PostRouter struct {
	Api *PostApi
}

func NewPostRouter(Api *PostApi) *PostRouter {
	return &PostRouter{Api}
}

func (s *PostRouter) RegisterRouter(Router *gin.RouterGroup) {
	router := Router.Group("post")
	{
		router.GET("getList", s.Api.GetList)
		router.POST("save", s.Api.Save)
		router.POST("delByIds", s.Api.DelByIds)
		router.POST("update", s.Api.Update)
	}
}
