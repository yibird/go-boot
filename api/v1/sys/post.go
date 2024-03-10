package sys

import (
	"github.com/gin-gonic/gin"
	res "go-boot/model/common/response"
	"go-boot/model/sys"
	"go-boot/utils/slice"
)

type PostApi struct {
}

// GetList
// @Tags 	  Post
// @Summary   获取岗位列表
// @accept    application/json
// @Produce   application/json
// @Success   200   {object} response.Response{data=object,msg=string}  "获取岗位列表"
// @Router    /post/getList [get]
func (s *PostApi) GetList(c *gin.Context) {
	var postQuery sys.PostQuery
	if err := c.ShouldBindQuery(&postQuery); err != nil {
		return
	}
	pageResult, err := postService.GetList(postQuery)
	res.ApiResultWithData(c, err, pageResult)
}

// Save 新增岗位
// @Tags Post
// @Summary   新增岗位
// @Produce   application/json
// @Success   200   {object} response.Response{}  "新增岗位"
// @Router    /post/save [post]
func (s *PostApi) Save(c *gin.Context) {
	post := sys.Post{}
	postDto := sys.PostDto{}
	res.SaveResult(c, postService.BaseService.Save(c, post, postDto))
}

// DelByIds Del 删除岗位
// @Tags Post
// @Summary   删除岗位
// @Produce   application/json
// @Success   200   {object} response.Response{}  "删除岗位"
// @Router    /post/save [post]
func (s *PostApi) DelByIds(c *gin.Context) {
	ids := slice.ToIntSlice(c.QueryArray("ids"))
	res.DelResult(c, postService.BaseService.DelByIds(&sys.Post{}, ids))
}

// Update 修改岗位
// @Tags Post
// @Summary   修改岗位
// @Produce   application/json
// @Success   200   {object} response.Response{}  "修改岗位"
// @Router    /post/save [post]
func (s *PostApi) Update(c *gin.Context) {
	post := sys.Post{}
	postDto := sys.PostDto{}
	res.UpdateResult(c, postService.BaseService.Update(c, post, postDto))
}
