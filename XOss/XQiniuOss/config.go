package XQiniuOss

type Config struct {
	AccessKey        string // 开发者key
	SecretKey        string // 开发者secret
	Bucket           string // 存储库名
	UseHTTPS         bool   // 是否使用https域名
	UseCdnDomains    bool   // 上传是否使用CDN上传加速
	UpTokenExpires   int    // 上传凭证有效期
	CallbackURL      string // 上传回调地址
	CallbackBodyType string // 上传回调信息格式
	EndUser          string // 唯一宿主标识
	FsizeMin         int    // 限定上传文件大小最小值，单位Byte。
	FsizeMax         int    // 限定上传文件大小最大值，单位Byte。超过限制上传文件大小的最大值会被判为上传失败，返回 413 状态码。
	MimeLimit        string // 限定上传类型
}

func DefaultConfig() *Config {
	return &Config{
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
