package XEcho

import "github.com/labstack/echo/v4"

var echoEngine *echo.Echo

func XEngine() *echo.Echo {
	return echoEngine
}
