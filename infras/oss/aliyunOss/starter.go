package aliyunOss

import (
	"GoWebScaffold/infras"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var aliyunOssClient *oss.Client

func Client() *oss.Client {
	infras.Check(aliyunOssClient)
	return aliyunOssClient
}

type Starter struct {
	infras.BaseStarter
	cfg *Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("AliyunOss", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	aliyunOssClient, err = NewClient(s.cfg)
	infras.FailHandler(err)
	sctx.Logger().Info("AliyunOss Setup Successful!")
}

func RunForTesting(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			"",
			60,
			60,
			false,
			false,
			"",
			"",
			"",
			"",
			"http://oss-cn-shenzhen.aliyuncs.com",
			false,
			"",
		}
	}

	aliyunOssClient, err = NewClient(config)
	return err
}
