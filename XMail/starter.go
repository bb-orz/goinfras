package XMail

import (
	"fmt"
	"github.com/bb-orz/goinfras"
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
	return "XMail"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Mail", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	s.cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v \n", define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	if s.cfg.NoAuth {
		mailDialer = NewNoAuthDialer(s.cfg.Server, s.cfg.Port)
	} else {
		mailDialer = NewAuthDialer(s.cfg.Server, s.cfg.User, s.cfg.Password, s.cfg.Port)
	}
	sctx.Logger().Info(s.Name(), goinfras.StepSetup, "Mail Dialer Setuped! ")
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(mailDialer)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().OK(s.Name(), goinfras.StepCheck, "Mail Dialer Setup Successful! ")
		return true
	}
	return false
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.ResourcesGroup }
