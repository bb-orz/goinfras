package XNats

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
	return "XNats"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Nats", &define)
		goinfras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print Nats Config:", zap.Any("NatsConfig", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	natsMQPool, err = NewPool(s.cfg, sctx.Logger())
	goinfras.ErrorHandler(err)
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(natsMQPool)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Nats Pool Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Nats Pool Setup Successful!", s.Name()))
	return true
}

func (s *starter) Stop() {
	natsMQPool.Close()
}
