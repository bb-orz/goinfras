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
		goinfras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	fmt.Printf("XAliyunSms Config: %v \n", *define)
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	aliyunSmsClient, err = NewAliyunSmsClient(s.cfg)
	goinfras.ErrorHandler(err)
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(aliyunSmsClient)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: AliyunSms Client Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: AliyunSms Client Setup Successful!", s.Name()))
	return true
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.ResourcesGroup }
