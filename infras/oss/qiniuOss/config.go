package qiniuOss

type qiniuOssConfig struct {
	Switch           bool   `yaml:"Switch"`
	AccessKey        string `yaml:"AccessKey"`
	SecretKey        string `yaml:"SecretKey"`
	Bucket           string `yaml:"Bucket"`
	UseHTTPS         bool   `yaml:"UseHTTPS"`
	UseCdnDomains    bool   `yaml:"UseCdnDomains"`
	UpTokenExpires   int    `yaml:"UpTokenExpires"`
	CallbackURL      string `yaml:"CallbackURL"`
	CallbackBodyType string `yaml:"CallbackBodyType"`
	EndUser          string `yaml:"EndUser"`
	FsizeMin         int    `yaml:"FsizeMin"`
	FsizeMax         int    `yaml:"FsizeLimit"`
	MimeLimit        string `yaml:"MimeLimit"`
}
