package aliyunOss

import (
	"GoWebScaffold/infras"
)

type AliyunOssStarter struct {
	infras.BaseStarter
}

func (s *AliyunOssStarter) Init(sctx *StarterContext) {
	client := Init(sctx.GetConfig())
	sctx.SetAliyunOSSClient(client)
}
