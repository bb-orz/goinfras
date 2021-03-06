package XRedisPubSub

type Config struct {
	DbHost      string // 主机地址
	DbPort      int    // 主机端口
	DbAuth      bool   // 权限认证开关
	DbPasswd    string // 权限密码
	MaxActive   int64  // 最大活动连接数，0为无限
	MaxIdle     int64  // 最大闲置连接数，0为无限
	IdleTimeout int64  // 闲置超时时间，0位无限
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
