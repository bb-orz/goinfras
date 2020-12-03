package XRedis

import (
	"GoWebScaffold/infras"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func RedisComponent() *redis.Pool {
	infras.Check(pool)
	return pool
}

func SetComponent(p *redis.Pool) {
	pool = p
}
