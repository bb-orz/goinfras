package XMongo

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"goinfras"
	"goinfras/infras"
)

type starter struct {
	BaseStarter
	cfg *Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	return starter
}

func (s *starter) Name() string {
	return "XMongo"
}

func (s *starter) Init(sctx *StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Mongodb", &define)
		infras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print Mongodb Config:", zap.Any("Mongodb", *define))
}

func (s *starter) Setup(sctx *StarterContext) {
	var err error
	client, err = NewClient(s.cfg)
	FailHandler(err)
}

func (s *starter) Check(sctx *StarterContext) bool {
	err := Check(client)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: MongoDB Client Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: MongoDB Client Setup Successful!", s.Name()))
	return true
}

func (s *starter) Stop() {
	_ = client.Disconnect(context.TODO())
}
