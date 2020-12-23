package XGocache

import "time"

type Config struct {
	DefaultExpiration time.Duration // 默认超时时间
	CleanupInterval   time.Duration // 内存清理间隔时间
}

func DefaultConfig() *Config {
	return &Config{
		time.Duration(3600) * time.Second,
		time.Duration(3600*24) * time.Second,
	}
}
