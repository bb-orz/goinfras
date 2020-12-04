package XQiniuOss

var qiniuOssClient *QnClient

func XClient() *QnClient {
	return qiniuOssClient
}

// 资源组件闭包执行
func XFClient(f func(c *QnClient) error) error {
	return f(qiniuOssClient)
}
