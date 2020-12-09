package XAliyunSms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	jsoniter "github.com/json-iterator/go"
)

type CommonSms struct {
	client *dysmsapi.Client
}

/**
 * @Description: 单条短信发送
 * @receiver c
 * @param scheme 交互协议："https"...
 * @param tel 手机号码
 * @param signName 必须，短信签名名称。请在控制台签名管理页面签名名称一列查看。必须是已添加、并通过审核的短信签名。
 * @param templateCode 必须，短信模板ID。请在控制台模板管理页面模板CODE一列查看。 必须是已添加、并通过审核的短信签名；且发送国际/港澳台消息时，请使用国际/港澳台短信模版。
 * @param templateParamJson 必须，短信模板参数，与templateCode 配合
 * @return *dysmsapi.SendSmsResponse
 * @return error
 */
func (c *CommonSms) SendSmsMsg(scheme, tel, signName, templateCode, templateParamsJson string) (*dysmsapi.SendSmsResponse, error) {

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = scheme
	request.PhoneNumbers = tel
	request.SignName = signName
	request.TemplateCode = templateCode
	request.TemplateParam = templateParamsJson
	return c.client.SendSms(request)
}

//

/**
 * @Description: 批量短信发送
 * @receiver c
 * @param scheme   		交互协议："https"...
 * @param templateCode 	必须，短信模板ID。请在控制台模板管理页面模板CODE一列查看。 必须是已添加、并通过审核的短信签名；且发送国际/港澳台消息时，请使用国际/港澳台短信模版。
 * @param tels			需要接收短信的手机号码数组
 * @param signNames		必须，短信签名名称数组。请在控制台签名管理页面签名名称一列查看。必须是已添加、并通过审核的短信签名。
 * @param templateParamsJson 必须，短信模板参数，与templateCode 配合
 * @return *dysmsapi.SendBatchSmsResponse
 * @return error
 */
func (c *CommonSms) SendBatchSmsMsg(scheme, templateCode, templateParamsJson string, tels, signNames []string) (*dysmsapi.SendBatchSmsResponse, error) {
	telJson, err := jsoniter.Marshal(tels)
	if err != nil {
		return nil, err
	}
	signNameJson, err := jsoniter.Marshal(signNames)
	if err != nil {
		return nil, err
	}
	request := dysmsapi.CreateSendBatchSmsRequest()
	request.Scheme = scheme
	request.PhoneNumberJson = string(telJson)
	request.SignNameJson = string(signNameJson)
	request.TemplateCode = templateCode
	request.TemplateParamJson = string(templateParamsJson)

	return c.client.SendBatchSms(request)
}
