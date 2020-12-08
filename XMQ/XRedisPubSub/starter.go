package XRedisPubSub

import (
	"fmt"
	"go.uber.org/zap"
	"goinfras"
)

type starter struct {
	goinfras.BaseStarter
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

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("RedisPubSub", &define)
		goinfras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print RedisPubSub Config:", zap.Any("RedisPubSubConfig", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	redisPubSubPool = NewRedisPubsubPool(s.cfg, sctx.Logger())
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(redisPubSubPool)
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
