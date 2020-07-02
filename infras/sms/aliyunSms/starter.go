package aliyunSms

import (
	"GoWebScaffold/infras"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/tietang/props/kvs"
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
	configs := sctx.Configs()
	define := AliyunSmsConfig{}
	err := kvs.Unmarshal(configs, &define, "AliyunSms")
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
		config = &AliyunSmsConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err = p.Unmarshal(config)
		if err != nil {
			return err
		}
	}
	aliyunSmsClient, err = NewAliyunSmsClient(config)
	return err
}
