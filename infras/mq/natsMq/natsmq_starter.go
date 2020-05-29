package natsMq

import (
	"GoWebScaffold/infras"
)

type NatsMQStarter struct {
	infras.BaseStarter
}

func (s *NatsMQStarter) Init(sctx *StarterContext) {
	natsMqPool := NatsMqPoolInit(sctx.GetConfig(), sctx.GetCommonLogger())
	sctx.SetNatsMQPool(natsMqPool)
}
