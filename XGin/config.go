package XGin

import "github.com/gin-gonic/gin"

type Config struct {
	Mode       string // 模式选择：debug
	ListenHost string // 服务运行ip
	ListenPort int    // 服务运行端口
	Tls        bool   // HTTPS相关配置，开关
	CertFile   string // HTTPS相关配置，证书文件
	KeyFile    string // HTTPS相关配置，私匙文件
}

// 默认最小启动配置
func DefaultConfig() *Config {
	return &Config{
		Mode:       gin.DebugMode,
		ListenHost: "127.0.0.1",
		ListenPort: 8090,
	}
}
