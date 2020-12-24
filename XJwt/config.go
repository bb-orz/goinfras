package XJwt

type Config struct {
	UseCache            bool   // 是否打开redis缓存
	PrivateKey          string // jwt编码私钥
	ExpSeconds          int64  // jwt编码超时时间
	TokenCacheKeyPrefix string // 缓存令牌键名前缀
}

func DefaultConfig() *Config {
	return &Config{
		UseCache:            true,
		PrivateKey:          "ginger_key",
		ExpSeconds:          6,
		TokenCacheKeyPrefix: "user.jwt.token.",
	}
}
