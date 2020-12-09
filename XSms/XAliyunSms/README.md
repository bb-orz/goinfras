# XAliyunSms Starter

> 基于 https://github.com/aliyun/alibaba-cloud-sdk-go 阿里云官方包
> 短信服务需引用 
> import "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

### AliyunSms Documentation

> Documentation https://help.aliyun.com/product/44282.html


### 短信使用流程：
- 入驻阿里云
- 开通短信服务
- 获取AccessKey
- 创建签名和模版
- 短信接口配置
- 发送短信

### XAliyunSms Starter Usage
```
goinfras.RegisterStarter(XAliyunSms.NewStarter())

```

### XAliyunSms Config Setting

```
Scheme          string // 交互协议："https"...
EndPoint        string // 必须，服务器节点
AccessKeyId     string // 必须，主账号AccessKey的ID。
AccessSecret    string // 必须，主账号秘钥。
SignName        string // 必须，短信签名名称。请在控制台签名管理页面签名名称一列查看。必须是已添加、并通过审核的短信签名。
TemplateCode    string // 必须，短信模板ID。请在控制台模板管理页面模板CODE一列查看。 必须是已添加、并通过审核的短信签名；且发送国际/港澳台消息时，请使用国际/港澳台短信模版。
Action          string // 系统规定参数。取值：SendSms。
OutId           string // 外部流水扩展字段。
SmsUpExtendCode string // 上行短信扩展码，无特殊需要此字段的用户请忽略此字段。

```

### XAliyunSms Usage

1、发送单条短信

```
response, err := XAliyunSms.XCommonSms().SendSmsMsg("scheme", "tel", "signName", "templateCode", "templateParamJson")
So(err, ShouldBeNil)
Println("Send Sms Status:", response.IsSuccess())
```

2、批量发送短信
```
smsResponse, err := XAliyunSms.XCommonSms().SendBatchSmsMsg("scheme", "templateCode","templateParamsJson", []string{"tel1","tel2"}, []string{"signName1","signName2"})
So(err, ShouldBeNil)
Println("Send Batch Sms Status:", smsResponse.IsSuccess())
```


