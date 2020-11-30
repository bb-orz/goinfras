package redisPubSub

import (
	"GoWebScaffold/infras"
	redigo "github.com/garyburd/redigo/redis"
)

var redisPubSubPool *redigo.Pool

func SetComponent(p *redigo.Pool) {
	redisPubSubPool = p
}

func RedisPubSubComponent() *redigo.Pool {
	infras.Check(redisPubSubPool)
	return redisPubSubPool
}

// 从Redis连接池获取一个连接
func getRedisConn() redigo.Conn {
	conn := RedisPubSubComponent().Get()
	return conn
}

// 从Redis连接池获取一个PubSub连接
func getRedisPubSubConn() *redigo.PubSubConn {
	conn := RedisPubSubComponent().Get()
	psConn := redigo.PubSubConn{Conn: conn}
	return &psConn
}
