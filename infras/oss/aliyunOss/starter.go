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
	cfg *aliyunOssConfig
}

func (s *AliyunOssStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := aliyunOssConfig{}
	err := kvs.Unmarshal(configs, &define, "AliyunOss")
	if err != nil {
		panic(err.Error())
	}
	s.cfg = &define
}

func (s *AliyunOssStarter) Setup(sctx *infras.StarterContext) {}

func (s *AliyunOssStarter) Start(sctx *infras.StarterContext) {
	aliyunOssClient = NewClient(s.cfg)
}

func (s *AliyunOssStarter) Stop(ctx *infras.StarterContext) {
}
