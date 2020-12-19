package XGin

import "github.com/gin-gonic/gin"

// 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO 错误处理

		c.Next()

	}
}
