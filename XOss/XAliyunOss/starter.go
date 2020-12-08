package XAliyunOss

import (
	"fmt"
	"go.uber.org/zap"
	"goinfras"
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
	return "XAliyunOss"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("AliyunOss", &define)
		goinfras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print AliyunOss Config:", zap.Any("AliyunOss", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	aliyunOssClient, err = NewClient(s.cfg)
	goinfras.ErrorHandler(err)
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(aliyunOssClient)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: AliyunOss Client Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: AliyunOss Client Setup Successful!", s.Name()))
	return true
}