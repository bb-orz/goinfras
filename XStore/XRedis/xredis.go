package XRedis

import (
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

var pool *redis.Pool

// 创建一个默认配置的DB
func CreateDefaultPool(config *Config, logger *zap.Logger) error {
	var err error
	if config == nil {
		config = DefaultConfig()
	}
	pool, err = NewPool(config, logger)
	return err
}
func XPool() *redis.Pool {
	return pool
}

// 资源组件闭包执行
func XFPool(f func(p *redis.Pool) error) error {
	return f(pool)
}

// Redis通用操作实例
func XCommon() *CommonRedisDao {
	dao := new(CommonRedisDao)
	dao.pool = XPool()
	return dao
}

// 检查连接池实例
func CheckPool() bool {
	if pool != nil {
		return true
	}
	return false
}
