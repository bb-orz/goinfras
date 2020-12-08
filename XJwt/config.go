package XJwt

type Config struct {
	PrivateKey string // jwt编码私钥
	ExpSeconds int64  // jwt编码超时时间
}

func DefaultConfig() *Config {
	return &Config{
		PrivateKey: "ginger_key",
		ExpSeconds: 6,
	}
}
