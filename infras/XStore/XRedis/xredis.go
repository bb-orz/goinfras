package XRedis

import (
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
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

/*实例化资源用于测试*/
func TestingInstantiation() error {
	var err error
	config := &Config{
		"127.0.0.1",
		6379,
		false,
		"",
		0,
		50,
		60,
	}

	pool, err = NewPool(config, zap.L())
	return err
}
