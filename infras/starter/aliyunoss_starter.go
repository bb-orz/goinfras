package starter

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/oss/aliyunOss"
)

type AliyunOssStarter struct {
	infras.BaseStarter
}

func (s *AliyunOssStarter) Init(sctx *StarterContext) {
	client := aliyunOss.Init(sctx.GetConfig())
	sctx.SetAliyunOSSClient(client)
}
