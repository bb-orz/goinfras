package qiniuOss

type qiniuOssConfig struct {
	AccessKey        string
	SecretKey        string
	Bucket           string
	UseHTTPS         bool
	UseCdnDomains    bool
	UpTokenExpires   int
	CallbackURL      string
	CallbackBodyType string
	EndUser          string
	FsizeMin         int
	FsizeMax         int
	MimeLimit        string
}
