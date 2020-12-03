package XAliyunSms

import (
	"GoWebScaffold/infras"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type Starter struct {
	infras.BaseStarter
	cfg Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("AliyunSms", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	var c *dysmsapi.Client
	c, err = NewAliyunSmsClient(&s.cfg)
	infras.FailHandler(err)
	SetCommponent(c)
}
