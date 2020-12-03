package XQiniuOss

import (
	"GoWebScaffold/infras"
	qiniu "github.com/qiniu/api.v7/v7/auth/qbox"
)

type Starter struct {
	infras.BaseStarter
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("QiniuOss", &define)
	infras.FailHandler(err)
	SetQnClient(define)
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var mac *qiniu.Mac
	mac = NewQiniuOssMac(&qiniuOssClient.cfg)
	SetMac(mac)
	sctx.Logger().Info("QiniuOss Setup Successful!")
}
