package core

import (
	"GoWebScaffold/core/verified"
	"GoWebScaffold/services"
	"sync"
)

// 服务层，实现services包定义的服务并设置该服务的实例，
// 需在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑

var _ services.ISmsService = new(SmsService)

func init() {
	// 初始化该业务模块时实例化服务
	var once sync.Once
	once.Do(func() {
		smsService := new(SmsService)
		services.SetSmsService(smsService)
	})
}

type SmsService struct {
	verifiedDomain *verified.VerifiedDomain
}

// 发送短信验证码
func (service *SmsService) SendPhoneVerifiedCode(dto services.SendPhoneVerifiedCodeDTO) error {
	panic("implement me")
}

// TODO 其他短信相关服务...
