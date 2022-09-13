package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 限流中间件，采用的是令牌桶算法

/**
  params: fillInterval是多长时间添加一个令牌（2*time.Second代表2秒添加一个令牌），cap是令牌桶的最大容量
		  当我们启动服务后，马上请求 http://127.0.0.1:8081/api/post/getPostListByPageByScore?page_size=3&page_number=1 能正常访问
		  但在第一次请求后两秒内再请求，会提示"你被限流了，请稍等"
  PS：若我们为所有路由都注册了限流中间件，它这个限流是将所有路由作为一个整体去看，其中一个路由被请求后2秒内另外一个路由被请求也会提示被限流了
*/
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	// 创建指定填充速率和容量大小的令牌桶
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		//if bucket.Take(1) > 0 {
		// 取令牌（非阻塞）
		if bucket.TakeAvailable(1) != 1 {
			//TakeAvailable()的返回值是令牌被取走的个数，若不等于1，则说明未能立即取到，需要等待
			c.String(http.StatusOK, "你被限流了，请稍等")
			c.Abort() // 不调用该请求的剩余处理程序
			return
		}
		// 取到令牌就放行
		c.Next() // 调用该请求的剩余处理程序
	}
}
