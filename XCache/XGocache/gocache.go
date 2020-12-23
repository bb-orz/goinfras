package XGocache

import (
	"github.com/pmylund/go-cache"
)

var goCache *cache.Cache

// 创建一个默认配置的DB
func CreateDefaultCache(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	goCache = NewCache(config)
}

// 检查连接池实例
func CheckPool() bool {
	if goCache != nil {
		return true
	}
	return false
}

func NewCache(cfg *Config) *cache.Cache {
	return cache.New(cfg.DefaultExpiration, cfg.CleanupInterval)
}
