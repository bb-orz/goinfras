package XEcho

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var echoEngine *echo.Echo

func NewEchoEngine(config *Config, middlewares ...echo.MiddlewareFunc) *echo.Echo {
	engine := echo.New()
	// Debug模式设置
	if config.Debug {
		engine.Debug = true
	}

	// 设置日志输出

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
