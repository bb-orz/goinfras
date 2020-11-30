package ginger

import (
	"GoWebScaffold/infras"
	"github.com/gin-gonic/gin"
)

/*资源组件化调用*/

var ginEngine *gin.Engine

// 设置组件资源
func SetComponent(engine *gin.Engine) {
	ginEngine = engine
}

// 组件化使用
func GinComponent() *gin.Engine {
	infras.Check(ginEngine)
	return ginEngine
}
