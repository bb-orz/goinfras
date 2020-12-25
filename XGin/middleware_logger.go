package XGin

import (
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func ZapLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		postForm := c.Request.PostForm.Encode()
		var reqBody []byte
		_, _ = c.Request.Body.Read(reqBody)

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		// 访问日志，记录响应时间、客户端请求信息等
		XLogger.XCommon().Info("[Global Request Log]",
			zap.String("ip", c.ClientIP()),
			zap.Duration("latency", latency),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("post", postForm),
			zap.String("body", string(reqBody)),
		)
	}
}
