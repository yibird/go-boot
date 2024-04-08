package sys

import (
	"github.com/gin-gonic/gin"
	"go-boot/core/model/res"
	"go-boot/utils/slice"
)

type PostApi struct {
	Service *PostService
}

func NewPostApi(Service *PostService) *PostApi {
	return &PostApi{Service}
}

// GetList
// @Tags 	  Post
// @Summary   获取岗位列表
// @accept    application/json
// @Produce   application/json
// @Success   200   {object} res.Response{data=object,msg=string}  "获取岗位列表"
// @Router    /post/getList [get]
func (s *PostApi) GetList(c *gin.Context) {
	var postQuery PostQuery
	if err := c.ShouldBindQuery(&postQuery); err != nil {
		return
	}
	pageResult, err := s.Service.GetList(postQuery)
	res.ApiResultWithData(c, err, pageResult)
}

// Save 新增岗位
// @Tags Post
// @Summary   新增岗位
// @Produce   application/json
// @Success   200   {object} res.Response{}  "新增岗位"
// @Router    /post/save [post]
func (s *PostApi) Save(c *gin.Context) {
	post := Post{}
	postDto := PostDto{}
	res.SaveResult(c, s.Service.BaseService.Save(c, post, postDto))
}

// DelByIds Del 删除岗位
// @Tags Post
// @Summary   删除岗位
// @Produce   application/json
// @Success   200   {object} res.Response{}  "删除岗位"
// @Router    /post/save [post]
func (s *PostApi) DelByIds(c *gin.Context) {
	ids := slice.ToIntSlice(c.QueryArray("ids"))
	res.DelResult(c, s.Service.BaseService.DelByIds(&Post{}, ids))
}

// Update 修改岗位
// @Tags Post
// @Summary   修改岗位
// @Produce   application/json
// @Success   200   {object} res.Response{}  "修改岗位"
// @Router    /post/save [post]
func (s *PostApi) Update(c *gin.Context) {
	post := Post{}
	postDto := PostDto{}
	res.UpdateResult(c, s.Service.BaseService.Update(c, post, postDto))
}
