package redisPubSub

type RedisPubSubConfig struct {
	Switch      bool   `val:"false"`
	DbHost      string `val:"127.0.0.1"`
	DbPort      int    `val:"6379"`
	DbAuth      bool   `val:"false"`
	DbPasswd    string
	MaxActive   int64 `val:"0"`
	MaxIdle     int64 `val:"50"`
	IdleTimeout int64 `val:"60"`
}
