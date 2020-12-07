package XNats

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
	return "XNats"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Nats", &define)
		infras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print Nats Config:", zap.Any("NatsConfig", *define))
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	var err error
	natsMQPool, err = NewPool(s.cfg, sctx.Logger())
	infras.FailHandler(err)
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(natsMQPool)
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
