package XRedis

import (
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func XPool() *redis.Pool {
	return pool
}

// 资源组件闭包执行
func XFDB(f func(p *redis.Pool) error) error {
	return f(pool)
}

// Redis通用操作实例
func XCommon() *CommonRedisDao {
	dao := new(CommonRedisDao)
	dao.pool = XPool()
	return dao
}
