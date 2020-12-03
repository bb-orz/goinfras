package XNats

import (
	"GoWebScaffold/infras"
)

type Starter struct {
	infras.BaseStarter
	cfg Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("NatsMq", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	natsMQPool, err = NewPool(&s.cfg, sctx.Logger())
	infras.FailHandler(err)
	sctx.Logger().Info("NatsMQPool Setup Successful!")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	natsMQPool.Close()
}
