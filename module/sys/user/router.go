package sys

import "github.com/gin-gonic/gin"

type UserRouter struct {
	Api *UserApi
}

func NewUserRouter(Api *UserApi) *UserRouter {
	return &UserRouter{Api}
}

func (s *UserRouter) RegisterRouter(Router *gin.RouterGroup) {
	router := Router.Group("user")
	{
		router.POST("save", s.Api.Save)
		router.GET("getList", s.Api.GetList)
	}
}
