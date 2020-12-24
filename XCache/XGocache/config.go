package XGocache

import "time"

type Config struct {
	DefaultExpiration time.Duration // 默认超时时间，小于1则长期有效
	CleanupInterval   time.Duration // 内存清理间隔时间，小于1则长期有效
	DumpFileName      string        // 缓存keys导出文件位置，停机时自动备份
}

func DefaultConfig() *Config {
	return &Config{
		time.Duration(3600) * time.Second,
		time.Duration(3600*24) * time.Second,
		"./key.json",
	}
}
