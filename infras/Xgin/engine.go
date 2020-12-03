package Xgin

import (
	"github.com/gin-gonic/gin"
)

func NewGinEngine(cfg *Config, middlewares ...gin.HandlerFunc) *gin.Engine {
	var engine *gin.Engine

	// 1.创建一个gin实例
	engine = gin.New()
	// 2.设置中间件
	engine.Use(middlewares...)

	return engine
}
