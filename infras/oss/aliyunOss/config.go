package aliyunOss

type aliyunOssConfig struct {
	Switch          bool   `yaml:"Switch"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	ConnTimeout     int    `yaml:"ConnTimeout"`
	RWTimeout       int    `yaml:"RwTimeout"`
	EnableMD5       bool   `yaml:"EnableMD5"`
	EnableCRC       bool   `yaml:"EnableCRC"`
	AuthProxy       string `yaml:"AuthProxy"`
	Proxy           string `yaml:"Proxy"`
	AccessKeyId     string `yaml:"AccessKeyId"`
	BucketName      string `yaml:"BucketName"`
	Endpoint        string `yaml:"Endpoint"`
	UseCname        bool   `yaml:"UseCname"`
	SecurityToken   string `yaml:"SecurityToken"`
}
