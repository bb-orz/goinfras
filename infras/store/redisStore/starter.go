package redisStore

import (
	"GoWebScaffold/infras"
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
)

var rPool *redis.Pool

func RedisPool() *redis.Pool {
	infras.Check(rPool)
	return rPool
}

type RedisStarter struct {
	infras.BaseStarter
	cfg *RedisConfig
}

func (s *RedisStarter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := RedisConfig{}
	err := viper.UnmarshalKey("Redis", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *RedisStarter) Setup(sctx *infras.StarterContext) {
	var err error
	rPool, err = NewRedisPool(s.cfg, sctx.Logger())
	infras.FailHandler(err)
	sctx.Logger().Info("RedisPool Setup Successful!")
}

func (s *RedisStarter) Stop(sctx *infras.StarterContext) {
	RedisPool().Close()
}

func RunForTesting() error {
	var err error
	config := &RedisConfig{
		"127.0.0.1",
		6379,
		false,
		"",
		0,
		50,
		60,
	}

	rPool, err = NewRedisPool(config, zap.L())
	return err
}
