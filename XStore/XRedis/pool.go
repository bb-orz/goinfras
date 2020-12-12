package XRedis

import (
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"strconv"
	"time"
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

// 检查连接池实例
func CheckPool() bool {
	if pool != nil {
		return true
	}
	return false
}

func NewPool(cfg *Config, logger *zap.Logger) (pool *redis.Pool, err error) {

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
				logger.Error("redis dial fatal", zap.Error(err))
				return nil, err
			}
			// 权限认证
			if cfg.DbAuth {
				if _, err := conn.Do("Auth", cfg.DbPasswd); err != nil {
					logger.Error("redis auth fatal", zap.Error(err))
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
				logger.Info("Redis Connect Ping Successful!")
			}
			return err
		},
	}

	// 一般启动后不关闭连接池
	return pool, nil
}
