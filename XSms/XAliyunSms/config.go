package XAliyunSms

type Config struct {
	EndPoint     string // 必须，服务器节点
	AccessKeyId  string // 必须，主账号AccessKey的ID。
	AccessSecret string // 必须，主账号秘钥。
}

func DefaultConfig() *Config {
	return &Config{
		"dysmsapi.aliyuncs.com",
		"",
		"",
	}
}
