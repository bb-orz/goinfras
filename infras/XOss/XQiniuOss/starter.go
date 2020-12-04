package XQiniuOss

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
	return "XQiniuOss"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("QiniuOss", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	qiniuOssClient = NewQnClient(&s.cfg)
	sctx.Logger().Info("QiniuOss Setup Successful!")
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(qiniuOssClient)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: QiniuOss Client Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: QiniuOss Client Setup Successful!", s.Name()))
	return true
}
