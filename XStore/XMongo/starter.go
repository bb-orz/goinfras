package XMongo

import (
	"context"
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
	return "XMongo"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Mongodb", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	s.cfg = &define
	sctx.Logger().SDebug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v \n", define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	client, err = NewClient(s.cfg)
	if sctx.PassError(s.Name(), goinfras.StepSetup, err) {
		sctx.Logger().SInfo(s.Name(), goinfras.StepSetup, fmt.Sprintf("MongoDB Client Setuped! \n"))
	}
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(client)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().SInfo(s.Name(), goinfras.StepCheck, fmt.Sprintf("MongoDB Client Setup Successful! \n"))
		return true
	}
	return false
}

func (s *starter) Stop() {
	_ = client.Disconnect(context.TODO())
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.ResourcesGroup }
