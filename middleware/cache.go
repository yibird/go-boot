package middleware

import "github.com/gin-gonic/gin"

// Cache cache中间件
func Cache(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
