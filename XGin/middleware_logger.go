package XGin

import (
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func ZapLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		postForm := c.Request.PostForm.Encode()
		userAgent := c.Request.UserAgent()

		// Debug
		// var jsonData map[string]interface{}
		// ct := c.Request.Header.Get("Content-Type")
		// if ct == "application/json" {
		// 	readCloser := c.Request.Body
		// 	decoder := json.NewDecoder(readCloser)
		// 	err := decoder.Decode(&jsonData)
		// 	if err != nil {
		// 		fmt.Printf("json error: %+v \n", err)
		// 	}
		// 	fmt.Printf("json data: %+v \n", jsonData)
		// }

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		zapFields := make([]zapcore.Field, 0)
		zapFields = append(zapFields,
			zap.String("IP", c.ClientIP()),
			zap.String("UserAgent", userAgent),
			zap.String("Method", c.Request.Method),
			zap.String("Path", path),
			zap.Int("Status", c.Writer.Status()),
			zap.Duration("Latency", latency),
		)
		if query != "" {
			zapFields = append(zapFields, zap.String("Query", query))
		}
		if postForm != "" {
			zapFields = append(zapFields, zap.String("PostForm", postForm))
		}

		// 访问日志，记录响应时间、客户端请求信息等
		XLogger.XCommon().Info("Global Request Log", zapFields...)
	}
}
