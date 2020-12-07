package XRedis

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
	return "XRedis"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Redis", &define)
		infras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print Redis Config:", zap.Any("Redis", *define))
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	var err error
	pool, err = NewPool(s.cfg, sctx.Logger())
	infras.FailHandler(err)
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(pool)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Redis Pool Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Redis Pool Setup Successful!", s.Name()))
	return true
}

func (s *starter) Stop() {
	pool.Close()
}
