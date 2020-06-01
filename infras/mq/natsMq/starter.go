package natsMq

import (
	"GoWebScaffold/infras"
	"github.com/tietang/props/kvs"
)

var natsMQPool *NatsPool

func NatsMQPool() *NatsPool {
	infras.Check(natsMQPool)
	return natsMQPool
}

type NatsMQStarter struct {
	infras.BaseStarter
	cfg *natsMqConfig
}

func (s *NatsMQStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := natsMqConfig{}
	err := kvs.Unmarshal(configs, &define, "NatsMq")
	infras.FailHandler(err)

	s.cfg = &define
}

func (s *NatsMQStarter) Setup(sctx *infras.StarterContext) {
	var err error
	natsMQPool, err = GetNatsMqPool(s.cfg, sctx.Logger())
	infras.FailHandler(err)
	sctx.Logger().Info("NatsMQPool Setup Successful!")
}

func (s *NatsMQStarter) Stop(sctx *infras.StarterContext) {
	NatsMQPool().Close()
}
