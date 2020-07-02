package aliyunSms

type AliyunSmsConfig struct {
	Scheme          string `val:"https"`
	EndPoint        string `val:"dysmsapi.aliyuncs.com"` // 必须，服务器节点
	AccessKeyId     string `val:""`                      // 必须，主账号AccessKey的ID。
	AccessSecret    string `val:""`                      // 必须，主账号秘钥。
	SignName        string `val:""`                      // 必须，短信签名名称。请在控制台签名管理页面签名名称一列查看。必须是已添加、并通过审核的短信签名。
	TemplateCode    string `val:""`                      // 必须，短信模板ID。请在控制台模板管理页面模板CODE一列查看。 必须是已添加、并通过审核的短信签名；且发送国际/港澳台消息时，请使用国际/港澳台短信模版。
	Action          string `val:"SendSms"`               // 系统规定参数。取值：SendSms。
	OutId           string `val:""`                      // 外部流水扩展字段。
	SmsUpExtendCode string `val:""`                      // 上行短信扩展码，无特殊需要此字段的用户请忽略此字段。
}
