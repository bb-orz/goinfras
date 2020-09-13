package aliyunOss

import (
	"GoWebScaffold/infras"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var aliyunOssClient *oss.Client

func AliyunOssClient() *oss.Client {
	infras.Check(aliyunOssClient)
	return aliyunOssClient
}

type AliyunOssStarter struct {
	infras.BaseStarter
	cfg *AliyunOssConfig
}

func (s *AliyunOssStarter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := AliyunOssConfig{}
	err := viper.UnmarshalKey("AliyunOss", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *AliyunOssStarter) Setup(sctx *infras.StarterContext) {
	var err error
	aliyunOssClient, err = NewClient(s.cfg)
	infras.FailHandler(err)
	sctx.Logger().Info("AliyunOss Setup Successful!")
}

func RunForTesting(config *AliyunOssConfig) error {
	var err error
	if config == nil {
		config = &AliyunOssConfig{
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
