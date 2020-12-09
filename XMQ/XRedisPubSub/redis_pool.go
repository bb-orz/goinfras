package XRedisPubSub

import (
	redigo "github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"strconv"
	"time"
)

/*
RedisPubSub用于实时性较高的消息推送，并不保证可靠,实现实时快速的消息异步分发功能。
使用场景：轻量级，高并发，延迟敏感，即时数据分析、秒杀计数器、缓存等，如果传输的数据量大不建议使用redis pubsub
Tips：原则上用于缓存的redis机器与用于pubsub的redis机器分开较好，如实在用同一个，只需在config配置填写一样即可。
*/

func NewRedisPubsubPool(cfg *Config, logger *zap.Logger) *redigo.Pool {
	// 配置并获得一个连接池对象的指针
	redisPubSubPool := &redigo.Pool{
		// 最大活动链接数。0为无限
		MaxActive: int(cfg.MaxActive),
		// 最大闲置链接数，0为无限
		MaxIdle: int(cfg.MaxIdle),
		// 闲置链接超时时间
		IdleTimeout: time.Duration(cfg.IdleTimeout) * time.Second,
		// 连接池的连接拨号
		Dial: func() (redigo.Conn, error) {
			// 连接
			redisAddr := cfg.DbHost + ":" + strconv.Itoa(cfg.DbPort)
			conn, err := redigo.Dial("tcp", redisAddr)
			if err != nil {
				logger.Error("redis dial fatal:", zap.Error(err))
				return nil, err
			}
			// 权限认证
			if cfg.DbAuth {
				if _, err := conn.Do("Auth", cfg.DbPasswd); err != nil {
					logger.Error("redis auth fatal:", zap.Error(err))
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},

		// 定时检测连接是否可用
		TestOnBorrow: func(conn redigo.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("Ping")
			if err != nil {
				logger.Warn("Redis PubSub Server Disconnect")
			}
			return err
		},
	}

	return redisPubSubPool
}
