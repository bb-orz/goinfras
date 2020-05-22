package starter

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/mq/redisPubSub"
)

type RedisPubSubStarter struct {
	infras.BaseStarter
}

func (s *RedisPubSubStarter) Init(sctx *StarterContext) {
	redisPS := redisPubSub.RedisMqInit(sctx.GetConfig(), sctx.GetCommonLogger())
	sctx.SetRedisPubSubMPool(redisPS)
}
