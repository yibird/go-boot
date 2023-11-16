package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS          = http.StatusOK
	BAD_REQUEST      = http.StatusBadRequest
	UNAUTHORIZED     = http.StatusUnauthorized
	FORBIDDEN        = http.StatusForbidden
	NOTFOUND         = http.StatusNotFound
	TOO_MANY_REQUEST = http.StatusTooManyRequests
	ERROR            = http.StatusInternalServerError
)

var StatusCodeMap = map[int]string{
	SUCCESS:          "OK",
	BAD_REQUEST:      "Bad Request",
	UNAUTHORIZED:     "Unauthorized",
	FORBIDDEN:        "Forbidden",
	NOTFOUND:         "Not Found",
	TOO_MANY_REQUEST: "Too Many Requests",
	ERROR:            "Internal Server Error",
}

func getStatus(code int) (int, string) {
	message, exists := StatusCodeMap[code]
	if exists {
		return code, message
	}
	return code, ""
}

// Response 统一响应结构体
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func ResultWithCode(statusCode int, c *gin.Context) {
	code, message := getStatus(statusCode)
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: nil,
		Msg:  message,
	})
}

func OK(c *gin.Context) {
	code, message := getStatus(SUCCESS)
	Result(code, nil, message, c)
}
func OkWithData(data interface{}, c *gin.Context) {
	code, message := getStatus(SUCCESS)
	Result(code, data, message, c)
}
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}
func OKWithRecord(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

// ----------------------------------

func Error(c *gin.Context) {
	code, message := getStatus(ERROR)
	Result(code, nil, message, c)
}
func ErrorWithData(data interface{}, c *gin.Context) {
	code, message := getStatus(ERROR)
	Result(code, data, message, c)
}
func ErrorWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}
func ErrorWithRecord(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

// ApiResultWithData ----------------------------------
func ApiResultWithData(c *gin.Context, err error, data interface{}) {
	if err == nil {
		OkWithData(data, c)
		return
	}
	Error(c)
}

func ApiResultWithMsg(c *gin.Context, err error, message string) {
	if err == nil {
		OkWithMessage(message, c)
		return
	}
	Error(c)
}
