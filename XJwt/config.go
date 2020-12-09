package XJwt

type Config struct {
	PrivateKey string // jwt编码私钥
	ExpSeconds int64  // jwt编码超时时间
	UseCache   bool   // 是否打开redis缓存
}

func DefaultConfig() *Config {
	return &Config{
		UseCache:   true,
		PrivateKey: "ginger_key",
		ExpSeconds: 6,
	}
}
