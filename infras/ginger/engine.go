package ginger

import (
	"github.com/gin-gonic/gin"
)

func NewGinEngine(cfg *ginConfig, middlewares ...gin.HandlerFunc) *gin.Engine {
	var engine *gin.Engine

	// 2.创建一个gin实例
	engine = gin.New()
	// 3.设置中间件
	engine.Use(middlewares...)

	return engine
}
