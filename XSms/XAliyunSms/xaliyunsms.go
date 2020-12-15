package XAliyunSms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func XClient() *dysmsapi.Client {
	return aliyunSmsClient
}

// 资源组件闭包执行
func XFClient(f func(c *dysmsapi.Client) error) error {
	return f(aliyunSmsClient)
}

func XCommonSms() *CommonSms {
	c := new(CommonSms)
	c.client = XClient()
	return c
}
