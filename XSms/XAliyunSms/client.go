package XAliyunSms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

var aliyunSmsClient *dysmsapi.Client

// 创建一个默认配置的Manager
func CreateDefaultClient(config *Config) error {
	var err error
	if config == nil {
		config = DefaultConfig()
	}
	aliyunSmsClient, err = NewAliyunSmsClient(config)
	return err
}

func NewAliyunSmsClient(config *Config) (*dysmsapi.Client, error) {
	return dysmsapi.NewClientWithAccessKey(config.EndPoint, config.AccessKeyId, config.AccessSecret)
}
