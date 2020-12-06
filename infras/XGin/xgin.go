package XGin

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var ginEngine *gin.Engine

func CreateDefaultEngine(config *Config) {
	// 1.配置gin中间件
	middlewares := make([]gin.HandlerFunc, 0)
	middlewares = append(middlewares, ZapLoggerMiddleware(zap.L()), ZapRecoveryMiddleware(zap.L(), false))

	if config == nil {
		config = DefaultConfig()
	}
	ginEngine = NewGinEngine(config, middlewares...)
}

// 资源组件实例调用
func XEngine() *gin.Engine {
	return ginEngine
}

// 资源组件闭包执行
func XFEngine(f func(c *gin.Engine) error) error {
	return f(ginEngine)
}
