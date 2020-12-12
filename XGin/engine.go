package XGin

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var ginEngine *gin.Engine

func NewGinEngine(cfg *Config, middlewares ...gin.HandlerFunc) *gin.Engine {
	var engine *gin.Engine

	// 1.创建一个gin实例
	engine = gin.New()
	// 2.设置中间件
	engine.Use(middlewares...)

	return engine
}

func CreateDefaultEngine(config *Config, logger *zap.Logger) {
	// 1.配置gin中间件
	middlewares := make([]gin.HandlerFunc, 0)
	middlewares = append(middlewares, ZapLoggerMiddleware(logger), ZapRecoveryMiddleware(logger, false))

	if config == nil {
		config = DefaultConfig()
	}
	ginEngine = NewGinEngine(config, middlewares...)
}
