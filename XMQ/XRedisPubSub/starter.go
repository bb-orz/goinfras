package XRedisPubSub

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
	return "XRedisPubSub"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("RedisPubSub", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	s.cfg = &define
	sctx.Logger().SDebug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v \n", define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	redisPubSubPool = NewRedisPubsubPool(s.cfg)
	sctx.Logger().SInfo(s.Name(), goinfras.StepSetup, fmt.Sprintf("RedisPubSub Pool Setuped!  \n"))
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	var err error
	err = goinfras.Check(redisPubSubPool)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		conn := redisPubSubPool.Get()
		defer conn.Close()
		_, err = conn.Do("PING", "ping")
		if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
			sctx.Logger().SInfo(s.Name(), goinfras.StepCheck, fmt.Sprintf("RedisPubSub Pool Setup Successful! \n"))
			return true
		}
	}
	return false
}

func (s *starter) Stop() {
	_ = redisPubSubPool.Close()
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.ResourcesGroup }
