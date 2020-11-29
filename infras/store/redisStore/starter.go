package redisStore

import (
	"GoWebScaffold/infras"
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
)

var pool *redis.Pool

func Pool() *redis.Pool {
	infras.Check(pool)
	return pool
}

type Starter struct {
	infras.BaseStarter
	cfg Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Redis", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	pool, err = NewPool(&s.cfg, sctx.Logger())
	infras.FailHandler(err)
	sctx.Logger().Info("RedisPool Setup Successful!")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	Pool().Close()
}

func RunForTesting() error {
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
