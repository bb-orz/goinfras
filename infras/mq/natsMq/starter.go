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
	if err != nil {
		panic(err.Error())
	}
	s.cfg = &define
}

func (s *NatsMQStarter) SetUp(sctx *infras.StarterContext) {}

func (s *NatsMQStarter) Start(sctx *infras.StarterContext) {
	natsMQPool = GetNatsMqPool(s.cfg)
}

func (s *NatsMQStarter) Stop(sctx *infras.StarterContext) {

}
