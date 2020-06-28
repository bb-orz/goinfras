package redisStore

import (
	"GoWebScaffold/infras"
	"github.com/garyburd/redigo/redis"
	"github.com/tietang/props/kvs"
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
	configs := sctx.Configs()
	define := RedisConfig{}
	err := kvs.Unmarshal(configs, &define, "Redis")
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
	config := RedisConfig{}
	p := kvs.NewEmptyCompositeConfigSource()
	err = p.Unmarshal(&config)
	if err != nil {
		return err
	}
	rPool, err = NewRedisPool(&config, zap.L())
	return err
}
