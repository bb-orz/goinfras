package XEcho

import (
	"github.com/labstack/echo/v4"
)

var echoEngine *echo.Echo

func NewEchoEngine(config *Config) *echo.Echo {
	engine := echo.New()
	// Debug模式设置
	if config.Debug {
		engine.Debug = true
	}

	// TODO 自定义engine

	return engine
}

func CreateDefaultEngine(config *Config) {
	// 1.配置gin中间件
	middlewares := make([]echo.MiddlewareFunc, 0)
	middlewares = append(middlewares, LoggerMiddleware(), RecoveryMiddleware(false))

	if config == nil {
		config = DefaultConfig()
	}
	echoEngine = NewEchoEngine(config)
	echoEngine.Use(middlewares...)
}
