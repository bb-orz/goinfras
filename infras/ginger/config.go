package ginger

type ginConfig struct {
	ListenHost string `val:"127.0.0.1"` // 服务运行ip
	ListenPort int    `val:"8090"`      // 服务运行端口
	cors       *corsConfig
	tls        bool   // HTTPS相关配置，开关
	certFile   string // HTTPS相关配置，证书文件
	keyFile    string // HTTPS相关配置，私匙文件
}

// Cors配置
type corsConfig struct {
	AllowAllOrigins  bool
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string
	MaxAge           int
	AllowOrigins     []string
	AllowMethods     []string
}
