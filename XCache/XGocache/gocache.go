package XGocache

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pmylund/go-cache"
	"os"
)

var goCache *cache.Cache

// 检查连接池实例
func CheckGocache() bool {
	if goCache != nil {
		return true
	}
	return false
}

func NewCache(cfg *Config) *cache.Cache {
	return cache.New(cfg.DefaultExpiration, cfg.CleanupInterval)
}

// 创建一个默认配置的DB
func CreateDefaultCache(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	goCache = NewCache(config)
}

// 创建一个默认配置的DB
func CreateDefaultCacheFrom(config *Config) error {
	var err error
	if config == nil {
		config = DefaultConfig()
	}
	goCache, err = NewCacheForm(config)
	if err != nil {
		return err
	}
	return nil
}

// 从文件读取
func NewCacheForm(cfg *Config) (*cache.Cache, error) {
	dumpFileName := cfg.DumpFileName
	dumpFile, err := os.OpenFile(dumpFileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	decoder := jsoniter.NewDecoder(dumpFile)
	var dumpData map[string]cache.Item
	err = decoder.Decode(&dumpData)
	if err != nil {
		return nil, err
	}
	cacheInstance := cache.NewFrom(cfg.DefaultExpiration, cfg.CleanupInterval, dumpData)
	return cacheInstance, nil
}

// 导出所有缓存keys，以便临时保存到文件或其他存储服务
func DumpItems(cfg *Config) error {
	items := goCache.Items()
	dumpFileName := cfg.DumpFileName
	dumpFile, err := os.OpenFile(dumpFileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	encoder := jsoniter.NewEncoder(dumpFile)
	return encoder.Encode(items)
}
