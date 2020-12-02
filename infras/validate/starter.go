package validate

import (
	"GoWebScaffold/infras"
	"go.uber.org/zap"
)

type Starter struct {
	infras.BaseStarter
	cfg Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Validate", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	if s.cfg.TransZh {
		validate, translator, err = NewZhValidator()
	} else {
		validate = NewValidator()
	}
	if err != nil {
		sctx.Logger().Error("Validator Error:", zap.Error(err))
	}
}
