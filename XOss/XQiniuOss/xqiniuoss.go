package XQiniuOss

import "goinfras"

var qiniuOssClient *QnClient

// 创建一个默认配置的Manager
func CreateDefaultClient(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	qiniuOssClient = NewQnClient(config)
}

func XClient() *QnClient {
	return qiniuOssClient
}

// 资源组件闭包执行
func XFClient(f func(c *QnClient) error) error {
	return f(qiniuOssClient)
}
