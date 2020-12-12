package XEcho

import "github.com/labstack/echo/v4"

func NewEchoEngine(confg *Config, middlewares ...echo.MiddlewareFunc) *echo.Echo {
	engine := echo.New()
	engine.Use(middlewares...)
	return engine
}
