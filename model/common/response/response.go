package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ApiResponse 统一响应结构体
type ApiResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = http.StatusOK
	ERROR   = http.StatusInternalServerError
	FAIL    = http.StatusGone
)

const (
	SuccessText = "响应成功"
	ErrorText   = "响应异常"
	FailText    = "响应失败"
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, ApiResponse{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, SuccessText, c)
}
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, SuccessText, c)
}
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}
func OKWithRecord(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

// ----------------------------------

func Fail(c *gin.Context) {
	Result(FAIL, map[string]interface{}{}, FailText, c)
}
func FailWithData(data interface{}, c *gin.Context) {
	Result(FAIL, data, FailText, c)
}
func FailWithMessage(message string, c *gin.Context) {
	Result(FAIL, map[string]interface{}{}, message, c)
}
func FailWithRecord(data interface{}, message string, c *gin.Context) {
	Result(FAIL, data, message, c)
}
