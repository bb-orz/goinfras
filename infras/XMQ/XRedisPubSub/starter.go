package XRedisPubSub

import (
	"GoWebScaffold/infras"
	"fmt"
)

type starter struct {
	infras.BaseStarter
	cfg Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = Config{}
	return starter
}

func (s *starter) Name() string {
	return "XRedisPubSub"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("RedisPubSub", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	redisPubSubPool = NewRedisPubsubPool(&s.cfg, sctx.Logger())
	sctx.Logger().Info("RedisPubSubPool Setup Successful ...")
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(redisPubSubPool)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: RedisPubSub Pool Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: RedisPubSub Pool Setup Successful!", s.Name()))
	return true
}

func (s *starter) Stop(sctx *infras.StarterContext) {
	_ = redisPubSubPool.Close()
}
