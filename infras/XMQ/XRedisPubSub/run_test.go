package XRedisPubSub

import (
	redigo "github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	var pool *redigo.Pool
	if config == nil {
		config = &Config{
			true,
			"127.0.0.1",
			6380,
			false,
			"",
			0,
			50,
			60,
		}

	}

	pool = NewRedisPubsubPool(config, zap.L())
	SetComponent(pool)
	return err
}
