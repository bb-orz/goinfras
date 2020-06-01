package qiniuOss

import (
	"GoWebScaffold/infras"
	qiniuOss "github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/tietang/props/kvs"
)

var qiniuOssClient *QnClient

type QnClient struct {
	mac *qiniuOss.Mac
	cfg *qiniuOssConfig
}

func (client *QnClient) GetMac() *qiniuOss.Mac {
	return client.mac
}

// 对外暴露的客户端
func QiniuOssClient() *QnClient {
	infras.Check(qiniuOssClient)
	return qiniuOssClient
}

type QiniuOssStarter struct {
	infras.BaseStarter
}

func (s *QiniuOssStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := qiniuOssConfig{}
	err := kvs.Unmarshal(configs, &define, "QiniuOss")
	infras.FailHandler(err)

	qiniuOssClient = new(QnClient)
	qiniuOssClient.cfg = &define

}

func (s *QiniuOssStarter) Setup(sctx *infras.StarterContext) {
	qiniuOssClient.mac = NewQiniuOssMac(qiniuOssClient.cfg)
	sctx.Logger().Info("QiniuOss Setup Successful!")
}
