package ginger

type Config struct {
	GinConfig
	CorsConfig
}

type GinConfig struct {
	ListenHost string // 服务运行ip
	ListenPort int    // 服务运行端口
	Tls        bool   // HTTPS相关配置，开关
	CertFile   string // HTTPS相关配置，证书文件
	KeyFile    string // HTTPS相关配置，私匙文件
}

// Cors配置
type CorsConfig struct {
	AllowAllOrigins  bool     // 是否允许所有源
	AllowHeaders     []string // 设置允许的头信息列表
	AllowCredentials bool     // 请求是否可以包括用户凭据，如cookies、HTTP身份验证或客户端SSL证书。
	ExposeHeaders    []string // 指定那些header项可以安全的导出
	MaxAge           int      // 指定预检前（option请求）请求的结果可以缓存多长时间（以秒为单位）
	AllowOrigins     []string // 设置允许的主机源列表
	AllowMethods     []string // 设置允许的请求方法列表
}
