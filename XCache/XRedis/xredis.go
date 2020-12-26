package XRedis

import (
	"github.com/gomodule/redigo/redis"
)

func XPool() *redis.Pool {
	return pool
}

// 资源组件闭包执行
func XFPool(f func(p *redis.Pool) error) error {
	return f(pool)
}

// Redis命令封装操作实例
func XCommand() *RedisCommand {
	dao := new(RedisCommand)
	dao.pool = XPool()
	return dao
}
