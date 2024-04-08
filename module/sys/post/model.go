package sys

import (
	"go-boot/core/model"
	"go-boot/core/model/req"
)

type Post struct {
	model.BaseModel
	ID        int64  `gorm:"primary" json:"id" mapstructure:"id"`
	PostCode  string `gorm:"post_code" json:"postCode" mapstructure:"postCode"`
	PostName  string `gorm:"post_name" json:"postName" mapstructure:"postName"`
	PostType  uint8  `gorm:"post_type" json:"postType" mapstructure:"postType"`
	PostLevel uint8  `gorm:"post_level" json:"postLevel" mapstructure:"postLevel"`
}

func (Post) TableName() string {
	return "sys_post"
}

type PostDto struct {
	ID        int64  `form:"id" json:"id"`
	PostCode  string `form:"postCode" json:"postCode"`
	PostName  string `form:"postName" json:"postName"`
	PostType  uint8  `form:"postType" json:"postType"`
	PostLevel uint8  `form:"postLevel" json:"postLevel"`
}

type PostQuery struct {
	req.BaseQueryModel
	PostCode string `form:"postCode" json:"postCode"`
	PostName string `form:"postName" json:"postName"`
}
