package XQiniuOss

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
	return "XQiniuOss"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("QiniuOss", &define)
		infras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print QiniuOss Config:", zap.Any("QiniuOss", *define))
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	qiniuOssClient = NewQnClient(s.cfg)
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
