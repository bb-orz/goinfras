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

func XClient() *dysmsapi.Client {
	return aliyunSmsClient
}

// 资源组件闭包执行
func XFClient(f func(c *dysmsapi.Client) error) error {
	return f(aliyunSmsClient)
}

func XCommonSms(config *Config) *CommonSms {
	c := new(CommonSms)
	c.client = XClient()
	c.cfg = config
	return c
}
