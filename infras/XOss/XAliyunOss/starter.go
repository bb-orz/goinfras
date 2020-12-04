package XAliyunOss

import (
	"GoWebScaffold/infras"
	"fmt"
)

type starter struct {
	infras.BaseStarter
	cfg Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = Config{}
	return starter
}

func (s *starter) Name() string {
	return "XAliyunOss"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("AliyunOss", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	var err error
	aliyunOssClient, err = NewClient(&s.cfg)
	infras.FailHandler(err)
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(aliyunOssClient)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: AliyunOss Client Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: AliyunOss Client Setup Successful!", s.Name()))
	return true
}
