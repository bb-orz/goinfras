package XNats

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
	return "XNats"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Nats", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	s.cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v ", define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	natsMQPool, err = NewPool(s.cfg)
	if sctx.PassError(s.Name(), goinfras.StepSetup, err) {
		sctx.Logger().Info(s.Name(), goinfras.StepSetup, "Nats Pool Setuped! ")
	}
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	var err error
	err = goinfras.Check(natsMQPool)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		_, err := natsMQPool.Get()
		if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
			sctx.Logger().OK(s.Name(), goinfras.StepCheck, "Nats Pool Setup Successful! ")
		}
		return true
	}
	return false
}

func (s *starter) Stop() {
	natsMQPool.Close()
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.ResourcesGroup }
