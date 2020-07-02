package aliyunSms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func NewAliyunSmsClient(config *AliyunSmsConfig) (*dysmsapi.Client, error) {
	return dysmsapi.NewClientWithAccessKey(config.EndPoint, config.AccessKeyId, config.AccessSecret)
}
