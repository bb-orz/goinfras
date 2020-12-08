package XValidate

import (
	"fmt"
	"go.uber.org/zap"
	"goinfras"
)

type starter struct {
	BaseStarter
	cfg *Config
}

func NewStarter() *starter {
	s := new(starter)
	s.cfg = &Config{}
	return s
}

func (s *starter) Name() string {
	return "XValidate"
}

func (s *starter) Init(sctx *StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Validate", &define)
		ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print Validate Config:", zap.Any("Validate", *define))
}

func (s *starter) Setup(sctx *StarterContext) {
	var err error
	if s.cfg.TransZh {
		validater, translater, err = NewZhValidater()
		ErrorHandler(err)
	} else {
		validater = NewValidater()
	}

}

func (s *starter) Check(sctx *StarterContext) bool {
	err1 := Check(validater)
	err2 := Check(translater)

	if err1 != nil || err2 != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Validater or Translater Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Validater and Translater Setup Successful!", s.Name()))
	return true
}
