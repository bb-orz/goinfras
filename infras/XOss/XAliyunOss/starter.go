package XAliyunOss

import (
	"GoWebScaffold/infras"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Starter struct {
	infras.BaseStarter
	cfg Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("AliyunOss", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	var c *oss.Client
	c, err = NewClient(&s.cfg)
	infras.FailHandler(err)
	SetComponent(c)
	sctx.Logger().Info("AliyunOss Setup Successful!")
}
