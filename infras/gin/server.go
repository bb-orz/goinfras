package gin

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func GinServer(cfg *ginConfig, logger *zap.Logger, middlewares ...gin.HandlerFunc) error {
	// 创建一个gin实例
	engine := gin.New()

	// 设置中间件
	engine.Use(middlewares...)

	return engine.Run(":" + strconv.Itoa(cfg.listenPort))
}
