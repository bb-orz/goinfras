package qiniuOss

import (
	"GoWebScaffold/infras"
)

type QiniuOssStarter struct {
	infras.BaseStarter
}

func (s *QiniuOssStarter) Init(sctx *StarterContext) {
	mac := Init(sctx.GetConfig())
	sctx.SetQiniuOSS(mac)
}
