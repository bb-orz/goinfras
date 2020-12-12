package XEcho

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewEchoEngine(confg *Config, middlewares ...echo.MiddlewareFunc) *echo.Echo {
	engine := echo.New()
	engine.Use(middlewares...)
	return engine
}

func CreateDefaultEngine(config *Config, logger *zap.Logger) {
	// 1.配置gin中间件
	middlewares := make([]echo.MiddlewareFunc, 0)
	middlewares = append(middlewares, LoggerMiddleware(logger), RecoveryMiddleware(logger, false))

	if config == nil {
		config = DefaultConfig()
	}
	echoEngine = NewEchoEngine(config, middlewares...)
}
