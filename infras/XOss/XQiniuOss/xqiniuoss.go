package XQiniuOss

var qiniuOssClient *QnClient

func XClient() *QnClient {
	return qiniuOssClient
}

// 资源组件闭包执行
func XFClient(f func(c *QnClient) error) error {
	return f(qiniuOssClient)
}

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
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
	qiniuOssClient = NewQnClient(config)
	return err
}
