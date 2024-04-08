package sys

import (
	"go-boot/core/model/res"
	"go-boot/core/service"
	"gorm.io/gorm"
)

type PostService struct {
	service.BaseService[Post, PostDto]
	DB *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{DB: db}
}

func (s *PostService) GetList(query PostQuery) (res.PageResult[Post], error) {
	db := s.DB.Where("post_code like ?", query.PostCode+"%").
		Or("post_name like ?", query.PostName+"%")
	page, err := res.SelectPage[Post](db, query.Current, query.PageSize)
	return page, err
}
