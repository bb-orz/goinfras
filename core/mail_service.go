package core

import (
	"GoWebScaffold/infras/validate"
	"GoWebScaffold/services"
	"sync"
)

// 服务层，实现services包定义的服务并设置该服务的实例，
// 需在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑

var _ services.IMailService = new(MailService)

func init() {
	// 初始化该业务模块时实例化服务
	var once sync.Once
	once.Do(func() {
		mailService := new(MailService)
		services.SetMailService(mailService)
	})
}

type MailService struct {
}

func (service *MailService) SendEmailForVerified(dto services.SendEmailForVerifiedDTO) error {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return err
	}

	return nil
}

func (service *MailService) SendEmailForgetPassword(dto services.SendEmailForgetPasswordDTO) error {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return err
	}

	return nil
}
