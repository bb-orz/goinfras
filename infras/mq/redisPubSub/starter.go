package redisPubSub

import (
	"GoWebScaffold/infras"
	redigo "github.com/garyburd/redigo/redis"
)

type Starter struct {
	infras.BaseStarter
	cfg Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("RedisPubSub", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var pool *redigo.Pool
	pool = NewRedisPubsubPool(&s.cfg, sctx.Logger())
	SetComponent(pool)
	sctx.Logger().Info("RedisPubSubPool Setup Successful ...")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	_ = RedisPubSubComponent().Close()
}
