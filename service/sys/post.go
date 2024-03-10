package sys

import (
	"go-boot/global"
	"go-boot/model/common/response"
	"go-boot/model/sys"
	"go-boot/service/common"
)

type PostService struct {
	common.BaseService[sys.Post, sys.PostDto]
}

var PostServiceApp = new(PostService)

func (s *PostService) GetList(query sys.PostQuery) (response.PageResult[sys.Post], error) {
	db := global.DB.Where("post_code like ?", query.PostCode+"%").
		Or("post_name like ?", query.PostName+"%")
	page, err := response.SelectPage[sys.Post](db, query.Current, query.PageSize)
	return page, err
}
