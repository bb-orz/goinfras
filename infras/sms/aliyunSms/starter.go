package aliyunSms

import (
	"GoWebScaffold/infras"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

var aliyunSmsClient *dysmsapi.Client
var sms *CommonSms

func Client() *dysmsapi.Client {
	infras.Check(aliyunSmsClient)
	return aliyunSmsClient
}

func Sms() *CommonSms {
	infras.Check(sms)
	return sms
}

type Starter struct {
	infras.BaseStarter
	cfg *Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("AliyunSms", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	aliyunSmsClient, err = NewAliyunSmsClient(s.cfg)
	infras.FailHandler(err)
	sms = NewCommonSms(s.cfg)
}

func RunForTesting(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
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
