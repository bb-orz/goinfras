package XAliyunSms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func NewAliyunSmsClient(config *Config) (*dysmsapi.Client, error) {
	return dysmsapi.NewClientWithAccessKey(config.EndPoint, config.AccessKeyId, config.AccessSecret)
}
