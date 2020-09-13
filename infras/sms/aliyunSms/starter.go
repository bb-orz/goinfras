package aliyunSms

import (
	"GoWebScaffold/infras"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

var aliyunSmsClient *dysmsapi.Client
var commonSms *CommonSms

func AliyunSmsClient() *dysmsapi.Client {
	infras.Check(aliyunSmsClient)
	return aliyunSmsClient
}

func AliyunCommonSms() *CommonSms {
	infras.Check(commonSms)
	return commonSms
}

type aliyunSmsStarter struct {
	infras.BaseStarter
	cfg *AliyunSmsConfig
}

func (s *aliyunSmsStarter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := AliyunSmsConfig{}
	err := viper.UnmarshalKey("AliyunSms", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *aliyunSmsStarter) Setup(sctx *infras.StarterContext) {
	var err error
	aliyunSmsClient, err = NewAliyunSmsClient(s.cfg)
	infras.FailHandler(err)
	commonSms = NewCommonSms(s.cfg)
}

func RunForTesting(config *AliyunSmsConfig) error {
	var err error
	if config == nil {
		config = &AliyunSmsConfig{
			"https",
			"dysmsapi.aliyuncs.com",
			"",
			"",
			"",
			"",
			"SendSms",
			"",
			"",
		}
	}
	aliyunSmsClient, err = NewAliyunSmsClient(config)
	return err
}
