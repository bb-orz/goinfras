package XValidate

import (
	"fmt"
	"github.com/bb-orz/goinfras"
)

type starter struct {
	goinfras.BaseStarter
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

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Validate", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	s.cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v ", define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	if s.cfg.TransZh {
		validater, translater, err = NewZhValidater()
		if sctx.PassError(s.Name(), goinfras.StepSetup, err) {
			sctx.Logger().Info(s.Name(), goinfras.StepSetup, "XValidate Utils Setuped! ")
		}
	} else {
		validater = NewValidater()
	}
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	var err error
	err = goinfras.Check(validater)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().OK(s.Name(), goinfras.StepCheck, "Validator Setup Successful! ")
	}
	err = goinfras.Check(translater)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().OK(s.Name(), goinfras.StepCheck, "Translator Setup Successful! ")
	}
	return true
}

func (s *starter) Stop() error {
	validater = nil
	return nil
}
