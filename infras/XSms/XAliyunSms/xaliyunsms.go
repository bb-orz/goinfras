package XAliyunSms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

var aliyunSmsClient *dysmsapi.Client

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
