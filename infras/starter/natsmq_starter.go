package starter

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/mq/natsMq"
)

type NatsMQStarter struct {
	infras.BaseStarter
}

func (s *NatsMQStarter) Init(sctx *StarterContext) {
	natsMqPool := natsMq.NatsMqPoolInit(sctx.GetConfig(), sctx.GetCommonLogger())
	sctx.SetNatsMQPool(natsMqPool)
}
