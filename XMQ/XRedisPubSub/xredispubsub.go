package XRedisPubSub

import (
	redigo "github.com/gomodule/redigo/redis"
)

// 资源组件实例调用
func XPool() *redigo.Pool {
	return redisPubSubPool
}

// 资源组件闭包执行
func XF(f func(c redigo.Conn) error) error {
	var err error
	var conn redigo.Conn
	// 获取连接
	conn = redisPubSubPool.Get()

	// 放回连接池
	defer func() {
		conn.Close()
	}()

	// 执行用户操作
	err = f(conn)
	if err != nil {
		return err
	}

	return nil
}

// 通用Publisher实例
func XRedisPublisher() *redisPublisher {
	publisher := new(redisPublisher)
	publisher.pool = XPool()
	return publisher
}

// 通用Subscriber实例
func XRedisSubscriber() *RedisSubscriber {
	subscriber := new(RedisSubscriber)
	subscriber.pool = XPool()
	return subscriber
}
