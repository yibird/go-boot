package sys

import "github.com/gin-gonic/gin"

type UserApi struct {
}

func NewUserApi() *UserApi {
	return &UserApi{}
}

func (s *UserApi) Save(c *gin.Context) {

}

func (s *UserApi) GetList(c *gin.Context) {

}
