package XAliyunOss

type Config struct {
	AccessKeyId     string //
	DefaultBucket   string // 存储库名
	AccessKeySecret string // 开发者AccessKeySecret
	Endpoint        string // 机房节点
	ConnTimeout     int    // 请求超时时间，包括连接超时、Socket读写超时，单位秒,默认连接超时30秒，读写超时60秒
	RWTimeout       int    // 读写超时设置
	EnableMD5       bool   // 是否开启MD5校验。推荐使用CRC校验，CRC的效率高于MD5
	EnableCRC       bool   // 是否开启CRC校验
	UseCname        bool   // 是否使用自定义域名CNAME
	AuthProxy       string // 带账号密码的代理服务器
	Proxy           string // 代理服务器，如http://8.8.8.8:3128
	SecurityToken   string // 临时用户的SecurityToken
}

func DefaultConfig() *Config {
	return &Config{
		"",
		"",
		"",
		"",
		60,
		60,
		false,
		false,
		false,
		"",
		"http://oss-cn-shenzhen.aliyuncs.com",
		"",
	}
}
