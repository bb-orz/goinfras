package RedisStore

// RedisDB配置
type redisConfig struct {
	DbHost      string `val:"127.0.0.1"`
	DbPort      int64  `val:"6379"`
	DbAuth      bool   `val:"false"`
	DbPasswd    string
	MaxActive   int64 `val:"0"`
	MaxIdle     int64 `val:"50"`
	IdleTimeout int64 `val:"60"`
}
