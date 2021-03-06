package XGin

import (
	"github.com/gin-gonic/gin"
)

// 资源组件实例调用
func XEngine() *gin.Engine {
	return ginEngine
}

// 资源组件闭包执行
func XFEngine(f func(c *gin.Engine) error) error {
	return f(ginEngine)
}
