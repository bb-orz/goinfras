package XAliyunSms

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
	return "XAliyunSms"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("AliyunSms", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().SDebug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %v \n", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	aliyunSmsClient, err = NewAliyunSmsClient(s.cfg)
	if sctx.PassError(s.Name(), goinfras.StepSetup, err) {
		sctx.Logger().SInfo(s.Name(), goinfras.StepSetup, fmt.Sprintf("Aliyun Sms Client Setuped! \n"))
	}
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(aliyunSmsClient)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().SInfo(s.Name(), goinfras.StepCheck, fmt.Sprintf("Aliyun Sms Client Setup Successful! \n"))
		return true
	}
	return false
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.ResourcesGroup }
