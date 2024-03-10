package sys

import "go-boot/global"

type User struct {
	global.BaseModel
}

func (User) TableName() string {
	return "sys_user"
}

type UserDto struct {
}

type UserQuery struct {
	global.BaseQueryModel
	UserName string `form:"username" json:"username"`
}
