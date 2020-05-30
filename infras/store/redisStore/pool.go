package redisStore

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

func NewRedisPool(cfg *redisConfig) (pool *redis.Pool, err error) {

	// 配置并获得一个连接池对象的指针
	pool = &redis.Pool{
		// 最大活动链接数。0为无限
		MaxActive: int(cfg.MaxActive),
		// 最大闲置链接数，0为无限
		MaxIdle: int(cfg.MaxIdle),
		// 闲置链接超时时间
		IdleTimeout: time.Duration(cfg.IdleTimeout) * time.Second,
		// 连接池的连接拨号
		Dial: func() (redis.Conn, error) {
			// 连接
			redisAddr := cfg.DbHost + ":" + strconv.Itoa(cfg.DbPort)
			conn, err := redis.Dial("tcp", redisAddr)
			if err != nil {
				fmt.Println("redis dial fatal:", err.Error())
				return nil, err
			}
			// 权限认证
			if cfg.DbAuth {
				if _, err := conn.Do("Auth", cfg.DbPasswd); err != nil {
					fmt.Println("redis auth fatal:", err.Error())
					fmt.Println(conn.Close().Error())
					return nil, err
				}
			}
			return conn, err
		},

		// 定时检测连接是否可用
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("Ping")
			if err != nil {
				fmt.Println("")
			}
			return err
		},
	}

	// 一般启动后不关闭连接池
	fmt.Println("Redis pool init ready!")
	return pool, nil
}
