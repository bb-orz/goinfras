package XGin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// Cors跨域请求处理中间件
func CORSMiddleware(cfg *CorsConfig) gin.HandlerFunc {
	if !cfg.AllowAllOrigins {
		return cors.New(cors.Config{
			AllowOrigins:     cfg.AllowOrigins,
			AllowMethods:     cfg.AllowMethods,
			AllowHeaders:     cfg.AllowHeaders,
			ExposeHeaders:    cfg.ExposeHeaders,
			AllowCredentials: cfg.AllowCredentials,
			MaxAge:           time.Second * time.Duration(cfg.MaxAge),
		})
	}

	return cors.Default()

}
