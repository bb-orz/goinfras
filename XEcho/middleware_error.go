package XEcho

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

/*错误处理中间件*/

func ErrorMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			err = next(c)

			// TODO 判断错误类型并处理错误

			return
		}
	}
}
