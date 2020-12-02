package aliyunSms

import (
	"GoWebScaffold/infras"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

var aliyunSmsClient *dysmsapi.Client

func SMSComponent() *dysmsapi.Client {
	infras.Check(aliyunSmsClient)
	return aliyunSmsClient
}

func SetCommponent(c *dysmsapi.Client) {
	aliyunSmsClient = c
}
