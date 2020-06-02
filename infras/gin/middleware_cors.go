package gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// Cors跨域请求处理中间件
func CORSMiddleware(cfg *corsConfig) gin.HandlerFunc {
	if cfg.AllowAllOrigins {

		// 允许所有跨域请求
		// same as
		// config := cors.DefaultConfig()
		// config.AllowAllOrigins = true
		return cors.Default()
	} else {
		// 如Request Header 无携带Origin字段，默认不是跨域请求CORS request
		return cors.New(cors.Config{
			AllowOrigins:     cfg.AllowOrigins,
			AllowMethods:     cfg.AllowMethods,
			AllowHeaders:     cfg.AllowHeaders,
			ExposeHeaders:    cfg.ExposeHeaders,
			AllowCredentials: cfg.AllowCredentials,
			MaxAge:           time.Second * time.Duration(cfg.MaxAge),
		})
	}
}
