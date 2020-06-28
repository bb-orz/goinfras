package aliyunOss

type AliyunOssConfig struct {
	AccessKeySecret string
	ConnTimeout     int
	RWTimeout       int
	EnableMD5       bool
	EnableCRC       bool
	AuthProxy       string
	Proxy           string
	AccessKeyId     string
	BucketName      string
	Endpoint        string `val:""http://oss-cn-shenzhen.aliyuncs.com"`
	UseCname        bool
	SecurityToken   string
}
