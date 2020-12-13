package XEcho

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

/* TODO 错误处理中间件 ,统一处理请求中的所有错误，需自定义错误类型 */

func ErrorMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			err = next(c)

			// TODO 判断错误类型并处理错误

			return
		}
	}
}
