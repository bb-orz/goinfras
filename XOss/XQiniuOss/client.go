package XQiniuOss

import (
	qiniuOss "github.com/qiniu/api.v7/v7/auth/qbox"
	"goinfras"
)

type QnClient struct {
	mac *qiniuOss.Mac
	cfg *Config
}

func NewQnClient(config *Config) *QnClient {
	client := new(QnClient)
	// 七牛云存储初始化
	client.mac = qiniuOss.NewMac(config.AccessKey, config.SecretKey)
	client.cfg = config
	return client
}
