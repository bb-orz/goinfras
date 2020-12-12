package XQiniuOss

import (
	qiniuOss "github.com/qiniu/api.v7/v7/auth/qbox"
)

var qiniuOssClient *QnClient

// 创建一个默认配置的Manager
func CreateDefaultClient(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	qiniuOssClient = NewQnClient(config)
}

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
