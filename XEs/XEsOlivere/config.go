package XEsOlivere

type Config struct {
	URL         string // 服务地址
	Username    string // 鉴权用户名
	Password    string // 鉴权密码
	Sniff       bool   // 启用或禁用嗅探器。
	Healthcheck bool   // 启用连接健康检查
	Infolog     string // Info级别日志记录文件路径
	Errorlog    string // Error级别日志记录文件路径
	Tracelog    string // Trace级别日志记录文件路径
}

func DefaultConfig() *Config {
	return &Config{}
}
