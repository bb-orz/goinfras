package ginger

type GinConfig struct {
	ListenHost string `val:"127.0.0.1"` // 服务运行ip
	ListenPort int    `val:"8090"`      // 服务运行端口
	Cors       *corsConfig
	Tls        bool   // HTTPS相关配置，开关
	CertFile   string // HTTPS相关配置，证书文件
	KeyFile    string // HTTPS相关配置，私匙文件
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
