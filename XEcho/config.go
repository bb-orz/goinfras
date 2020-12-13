package XEcho

import "github.com/labstack/echo/v4/middleware"

type Config struct {
	*EchoConfig
	*CorsConfig
}

type EchoConfig struct {
	Debug      bool   // 调试模式
	ListenHost string // 服务运行ip
	ListenPort int    // 服务运行端口
	Tls        bool   // HTTPS相关配置，开关
	CertFile   string // HTTPS相关配置，证书文件
	KeyFile    string // HTTPS相关配置，私匙文件
}

type CorsConfig struct {
	MainDomain       string             // 设置主域名
	AllowAllHost     bool               // 是否允许所有主机访问
	AllowSubDomain   bool               // 是否允许子域名访问
	Skipper          middleware.Skipper // 一些可以跳过该中间件的方法
	AllowOrigins     []string           `json:"allow_origins"`     // 定义允许的源
	AllowMethods     []string           `json:"allow_methods"`     // 定义允许请求的方法
	AllowHeaders     []string           `json:"allow_headers"`     // 定义允许的请求头
	AllowCredentials bool               `json:"allow_credentials"` // 是否允许打开证书
	ExposeHeaders    []string           `json:"expose_headers"`    // 定义被允许访问的白名单headers
	MaxAge           int                `json:"max_age"`           // 定义preflight请求时前请求的结果持续多长时间（以秒为单位）可以缓存。
}

// 默认最小启动配置
func DefaultConfig() *Config {
	return &Config{
		&EchoConfig{
			ListenHost: "127.0.0.1",
			ListenPort: 8090,
		},
		&CorsConfig{
			AllowAllHost: true,
		},
	}
}
