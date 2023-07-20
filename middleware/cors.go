package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors CORS跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求方法和请求源
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行预检查方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// CorsByRules CORS跨域中间件,根据规则进行跨域配置
func CorsByRules() gin.HandlerFunc {
	return func(c *gin.Context) {
		//// 获取请求方法和请求源
		//method := c.Request.Method
		//origin := c.Request.Header.Get("Origin")
	}
}
