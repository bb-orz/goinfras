package XNats

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
	return "XNats"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("NatsMq", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	var err error
	natsMQPool, err = NewPool(&s.cfg, sctx.Logger())
	infras.FailHandler(err)
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(natsMQPool)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Nats Pool Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Nats Pool Setup Successful!", s.Name()))
	return true
}

func (s *starter) Stop(sctx *infras.StarterContext) {
	natsMQPool.Close()
}
