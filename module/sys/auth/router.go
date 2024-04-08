package sys

import (
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	AuthApi *AuthApi
}

func NewAuthRouter(AuthApi *AuthApi) (*AuthRouter, error) {
	return &AuthRouter{AuthApi}, nil
}

func (s *AuthRouter) RegisterRouter(Router *gin.RouterGroup) {
	router := Router.Group("auth")
	api := s.AuthApi
	{
		router.GET("/getCaptcha", api.GetCaptcha)
		router.POST("/accountLogin", api.AccountLogin)
	}
}
