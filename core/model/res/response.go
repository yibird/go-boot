package res

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

	ADD_SUCCESS = 10001
	ADD_ERROR   = 10002
)
const (
	SUCCESS_MESSAGE     = "OK"
	BAD_REQUEST_MESSAGE = "Bad Request"
	SAVE_OK_MESSAGE     = "Save Successfully"
	SAVE_ERR_MESSAGE    = "Save Error"
	DEL_OK_MESSAGE      = "Delete Successfully"
	DEL_ERR_MESSAGE     = "Delete Error"
	UPDATE_OK_MESSAGE   = "Update Successfully"
	UPDATE_ERR_MESSAGE  = "Update Error"
)

var StatusCodeMap = map[int]string{
	SUCCESS:          SUCCESS_MESSAGE,
	BAD_REQUEST:      BAD_REQUEST_MESSAGE,
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
func OkWithStatus(c *gin.Context, status int) {
	code, message := getStatus(status)
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
func ErrorWithStatus(c *gin.Context, status int) {
	code, message := getStatus(status)
	Result(code, nil, message, c)
}
func ErrorWithMessage(message string, c *gin.Context) {
	Result(ERROR, nil, message, c)
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

func OkWithMsg(c *gin.Context, err error, message string) {
	if err == nil {
		OkWithMessage(message, c)
		return
	}
	Error(c)
}

func SaveResult(c *gin.Context, err error) {
	if err == nil {
		OkWithMessage(SAVE_OK_MESSAGE, c)
		return
	}
	ErrorWithMessage(SAVE_ERR_MESSAGE, c)
}
func DelResult(c *gin.Context, err error) {
	if err == nil {
		OkWithMessage(DEL_OK_MESSAGE, c)
		return
	}
	ErrorWithMessage(DEL_ERR_MESSAGE, c)
}
func UpdateResult(c *gin.Context, err error) {
	if err == nil {
		OkWithMessage(UPDATE_OK_MESSAGE, c)
		return
	}
	ErrorWithMessage(UPDATE_ERR_MESSAGE, c)
}
