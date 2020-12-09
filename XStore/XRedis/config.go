package XRedis

// RedisDB配置
type Config struct {
	DbHost      string // 主机地址
	DbPort      int    // 主机端口
	DbAuth      bool   // 是否开启鉴权
	DbPasswd    string // 鉴权密码
	MaxActive   int64  // 最大活动链接数。0为无限
	MaxIdle     int64  // 最大闲置链接数，0为无限
	IdleTimeout int64  // 闲置链接超时时间
}

func DefaultConfig() *Config {
	return &Config{
		"127.0.0.1",
		6379,
		false,
		"",
		0,
		50,
		60,
	}
}
