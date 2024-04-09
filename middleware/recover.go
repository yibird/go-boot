package middleware

import (
	"github.com/gin-gonic/gin"
	"go-boot/core/model/res"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				res.ErrorWithMessage(errorToStr(r), c)
				c.Abort()
			}
		}()
		c.Next()
	}
}

func errorToStr(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
