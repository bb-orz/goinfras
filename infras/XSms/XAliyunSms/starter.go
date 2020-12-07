package XAliyunSms

import (
	"GoWebScaffold/infras"
	"fmt"
	"go.uber.org/zap"
)

type starter struct {
	infras.BaseStarter
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

func (s *starter) Init(sctx *infras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("AliyunSms", &define)
		infras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print AliyunSms Config:", zap.Any("AliyunSms", *define))
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	var err error
	aliyunSmsClient, err = NewAliyunSmsClient(s.cfg)
	infras.FailHandler(err)
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(aliyunSmsClient)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: AliyunSms Client Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: AliyunSms Client Setup Successful!", s.Name()))
	return true
}
