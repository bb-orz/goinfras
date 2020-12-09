package XGlobal

type Config struct {
	Env        string // 允许环境：dev、testing、product
	AppName    string // 应用名称
	ServerName string // 服务名称
}

func DefaultConfig() *Config {
	return &Config{
		Env:        "1.0.0",
		AppName:    "Ginger App",
		ServerName: "My APP Host",
	}
}
