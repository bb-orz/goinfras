package XAliyunSms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	jsoniter "github.com/json-iterator/go"
)

type CommonSms struct {
	client *dysmsapi.Client
	cfg    *Config
}

func NewCommonSms() {
	c := new(CommonSms)
	c.client = SMSComponent()
}

func CommonCommonSms() *CommonSms {
	return new(CommonSms)
}

// 单条发送
func (c *CommonSms) SendSmsMsg(tel, code string) (*dysmsapi.SendSmsResponse, error) {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = c.cfg.Scheme
	request.PhoneNumbers = tel
	request.SignName = c.cfg.SignName
	request.TemplateCode = c.cfg.TemplateCode
	request.TemplateParam = "{\"code\":\"" + code + "\"}"
	return c.client.SendSms(request)
}

// 批量发送
func (c *CommonSms) SendBatchSmsMsg(tels, signNames []string, params []map[string]interface{}) (*dysmsapi.SendBatchSmsResponse, error) {
	telJson, err := jsoniter.Marshal(tels)
	if err != nil {
		return nil, err
	}
	signNameJson, err := jsoniter.Marshal(signNames)
	if err != nil {
		return nil, err
	}
	paramsJson, err := jsoniter.Marshal(params)
	if err != nil {
		return nil, err
	}
	request := dysmsapi.CreateSendBatchSmsRequest()
	request.Scheme = c.cfg.Scheme
	request.PhoneNumberJson = string(telJson)
	request.SignNameJson = string(signNameJson)
	request.TemplateCode = c.cfg.TemplateCode
	request.TemplateParamJson = string(paramsJson)

	return c.client.SendBatchSms(request)
}
