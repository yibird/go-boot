package sys

import (
	"go-boot/core/model"
	"go-boot/core/model/req"
)

type User struct {
	model.BaseModel
}

func (User) TableName() string {
	return "sys_user"
}

type UserDto struct {
}

type UserQuery struct {
	req.BaseQueryModel
	UserName string `form:"username" json:"username"`
}
