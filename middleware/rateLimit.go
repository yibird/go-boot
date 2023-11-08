package middleware

import (
	"go-boot/config"
	"go-boot/model/common/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// 限流中间件,基于令牌桶算法实现限流策略
// Parameters:
//   - limit: 每秒允许通过的请求速率。
//   - burst: 令牌桶的容量。
func RateLimit(ratelimit config.RateLimit) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(ratelimit.Limit), ratelimit.Burst)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
		} else {
			c.Abort()
			response.ResultWithCode(response.TOO_MANY_REQUEST, c)
		}
	}
}
