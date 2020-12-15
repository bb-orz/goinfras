package XEcho

type Config struct {
	Debug             bool   // 调试模式
	ListenHost        string // 服务运行ip
	ListenPort        int    // 服务运行端口
	Tls               bool   // HTTPS相关配置，开关
	CertFile          string // HTTPS相关配置，证书文件
	KeyFile           string // HTTPS相关配置，私匙文件
	UseSelfMiddleware bool   // 该启动器默认提供Logger、Recovery、Error三个初始化的中间件，如需重新自定义，设置为true
}

// 默认最小启动配置
func DefaultConfig() *Config {
	return &Config{
		Debug:      true,
		ListenHost: "127.0.0.1",
		ListenPort: 8090,
	}
}
