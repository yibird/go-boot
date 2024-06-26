package middleware

import (
	"github.com/gin-gonic/gin"
	"go-boot/config"
	"go-boot/core/model/res"
	"golang.org/x/time/rate"
)

// RateLimit 限流中间件
func RateLimit(ratelimit config.RateLimit) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(ratelimit.Limit), ratelimit.Burst)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
		} else {
			c.Abort()
			res.ResultWithCode(res.TOO_MANY_REQUEST, c)
		}
	}
}
