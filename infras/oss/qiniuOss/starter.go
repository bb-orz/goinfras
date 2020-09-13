package qiniuOss

import (
	"GoWebScaffold/infras"
	qiniuOss "github.com/qiniu/api.v7/v7/auth/qbox"
)

var qiniuOssClient *QnClient

type QnClient struct {
	mac *qiniuOss.Mac
	cfg *QiniuOssConfig
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
	viper := sctx.Configs()
	define := QiniuOssConfig{}
	err := viper.UnmarshalKey("QiniuOss", &define)
	infras.FailHandler(err)
	qiniuOssClient = new(QnClient)
	qiniuOssClient.cfg = &define

}

func (s *QiniuOssStarter) Setup(sctx *infras.StarterContext) {
	qiniuOssClient.mac = NewQiniuOssMac(qiniuOssClient.cfg)
	sctx.Logger().Info("QiniuOss Setup Successful!")
}

func RunForTesting(config *QiniuOssConfig) error {
	var err error
	if config == nil {
		config = &QiniuOssConfig{
			"",
			"",
			"",
			false,
			false,
			7200,
			"",
			"",
			"",
			1024,
			10485760,
			"",
		}
	}
	mac := NewQiniuOssMac(config)
	client := new(QnClient)
	client.mac = mac
	client.cfg = config
	qiniuOssClient = client
	return err
}
