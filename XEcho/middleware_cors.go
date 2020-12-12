package XEcho

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
	"strings"
)

// 允许所有请求的默认配置
var DefaultCORSConfig = &CorsConfig{
	Skipper:      middleware.DefaultSkipper,
	AllowOrigins: []string{"*"},
	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
}

// 默认CORS中间件调用方法
func DefaultCORSMiddleware() echo.MiddlewareFunc {
	return CORSMiddleware(DefaultCORSConfig)
}

// 用户配置的CORS中间件
func CORSMiddleware(config *CorsConfig) echo.MiddlewareFunc {
	// 一些配置项为空时，启用默认配置
	if config.Skipper == nil {
		config.Skipper = DefaultCORSConfig.Skipper
	}
	if len(config.AllowOrigins) == 0 {
		config.AllowOrigins = DefaultCORSConfig.AllowOrigins
	}
	if len(config.AllowMethods) == 0 {
		config.AllowMethods = DefaultCORSConfig.AllowMethods
	}

	allowMethods := strings.Join(config.AllowMethods, ",")
	allowHeaders := strings.Join(config.AllowHeaders, ",")
	exposeHeaders := strings.Join(config.ExposeHeaders, ",")
	maxAge := strconv.Itoa(config.MaxAge)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			origin := req.Header.Get(echo.HeaderOrigin)
			allowOrigin := ""

			// 检查合法的源
			for _, o := range config.AllowOrigins {
				if o == "*" || o == origin {
					allowOrigin = o
					break
				}
			}
			if config.AllowSubDomain && config.MainDomain != "" {
				if strings.Contains(origin, config.MainDomain) {
					allowOrigin = origin
				}
			}
			if config.AllowAllHost {
				allowOrigin = c.Scheme() + "://" + req.Host
			}

			// 简单请求 not options
			if req.Method != echo.OPTIONS {
				res.Header().Add(echo.HeaderVary, echo.HeaderOrigin)
				res.Header().Set(echo.HeaderAccessControlAllowOrigin, allowOrigin)
				if config.AllowCredentials {
					res.Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
				}
				if exposeHeaders != "" {
					res.Header().Set(echo.HeaderAccessControlExposeHeaders, exposeHeaders)
				}
				return next(c)
			}

			// options 预请求
			res.Header().Add(echo.HeaderVary, echo.HeaderOrigin)
			res.Header().Add(echo.HeaderVary, echo.HeaderAccessControlRequestMethod)
			res.Header().Add(echo.HeaderVary, echo.HeaderAccessControlRequestHeaders)
			res.Header().Set(echo.HeaderAccessControlAllowOrigin, allowOrigin)
			res.Header().Set(echo.HeaderAccessControlAllowMethods, allowMethods)
			if config.AllowCredentials {
				res.Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
			}
			if allowHeaders != "" {
				res.Header().Set(echo.HeaderAccessControlAllowHeaders, allowHeaders)
			} else {
				h := req.Header.Get(echo.HeaderAccessControlRequestHeaders)
				if h != "" {
					res.Header().Set(echo.HeaderAccessControlAllowHeaders, h)
				}
			}
			if config.MaxAge > 0 {
				res.Header().Set(echo.HeaderAccessControlMaxAge, maxAge)
			}
			return c.NoContent(http.StatusNoContent)
		}
	}
}
