package redisStore

import (
	"GoWebScaffold/infras"
	"github.com/garyburd/redigo/redis"
)

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
	var p *redis.Pool
	p, err = NewPool(&s.cfg, sctx.Logger())
	infras.FailHandler(err)
	SetComponent(p)
	sctx.Logger().Info("RedisPool Setup Successful!")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	RedisComponent().Close()
}
