package XRedisPubSub

import (
	"GoWebScaffold/infras"
	"fmt"
	"go.uber.org/zap"
)

type starter struct {
	infras.BaseStarter
	cfg *Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	return starter
}

func (s *starter) Name() string {
	return "XRedisPubSub"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("RedisPubSub", &define)
		infras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print RedisPubSub Config:", zap.Any("RedisPubSubConfig", *define))
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	redisPubSubPool = NewRedisPubsubPool(s.cfg, sctx.Logger())
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

func (s *starter) Stop() {
	_ = redisPubSubPool.Close()
}
