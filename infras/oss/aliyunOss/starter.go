package aliyunOss

import (
	"GoWebScaffold/infras"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/tietang/props/kvs"
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
	configs := sctx.Configs()
	define := AliyunOssConfig{}
	err := kvs.Unmarshal(configs, &define, "AliyunOss")
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
		config = &AliyunOssConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err = p.Unmarshal(config)
		if err != nil {
			return err
		}
	}

	aliyunOssClient, err = NewClient(config)
	return err
}
