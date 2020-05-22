package starter

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/oss/qiniuOss"
)

type QiniuOssStarter struct {
	infras.BaseStarter
}

func (s *QiniuOssStarter) Init(sctx *StarterContext) {
	mac := qiniuOss.Init(sctx.GetConfig())
	sctx.SetQiniuOSS(mac)
}
