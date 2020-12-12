package XEcho

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"time"
)

var timeFormat = "2019-11-09T23:02:28.844+0800"

func LoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			start := time.Now()
			path := c.Request().URL.Path
			query := c.Request().URL.RawQuery
			postForm := c.Request().PostForm.Encode()
			var reqBody []byte
			_, _ = c.Request().Body.Read(reqBody)

			// Next
			err = next(c)

			// 从上个中间价回来
			end := time.Now()
			latency := end.Sub(start)
			if err != nil {
				logger.Error(err.Error())
			} else {
				logger.Info("[Global Request Log]",
					zap.String("host", c.Request().Host),
					zap.Duration("latency", latency),
					zap.Int("status", c.Response().Status),
					zap.String("method", c.Request().Method),
					zap.String("path", path),
					zap.String("query", query),
					zap.String("post", postForm),
					zap.String("body", string(reqBody)),
				)
			}
			return
		}
	}
}
