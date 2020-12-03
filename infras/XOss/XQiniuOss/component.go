package XQiniuOss

import (
	"GoWebScaffold/infras"
	qiniu "github.com/qiniu/api.v7/v7/auth/qbox"
)

var qiniuOssClient *QnClient

type QnClient struct {
	mac *qiniu.Mac
	cfg Config
}

func QiniuOssComponent() *QnClient {
	return qiniuOssClient
}

func SetQnClient(cfg Config) {
	qiniuOssClient = new(QnClient)
	qiniuOssClient.cfg = cfg
}

func SetMac(m *qiniu.Mac) {
	client := new(QnClient)
	client.mac = m
}

// 对外暴露的客户端
func Client() *QnClient {
	infras.Check(qiniuOssClient)
	return qiniuOssClient
}
