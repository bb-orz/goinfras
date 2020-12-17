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
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Validate", &define)
		goinfras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	fmt.Printf("Print XValidate Config: %v \n", *define)
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	if s.cfg.TransZh {
		validater, translater, err = NewZhValidater()
		goinfras.ErrorHandler(err)
	} else {
		validater = NewValidater()
	}

}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err1 := goinfras.Check(validater)
	err2 := goinfras.Check(translater)

	if err1 != nil || err2 != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Validater or Translater Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Validater and Translater Setup Successful!", s.Name()))
	return true
}
