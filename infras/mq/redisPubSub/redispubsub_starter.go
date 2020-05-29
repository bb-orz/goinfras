package redisPubSub

import (
	"GoWebScaffold/infras"
)

type RedisPubSubStarter struct {
	infras.BaseStarter
}

func (s *RedisPubSubStarter) Init(sctx *StarterContext) {
	redisPS := RedisMqInit(sctx.GetConfig(), sctx.GetCommonLogger())
	sctx.SetRedisPubSubMPool(redisPS)
}
