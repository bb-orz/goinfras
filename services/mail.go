package services

import "GoWebScaffold/infras"

/* 定义邮件服务模块的服务层方法，并定义数据传输对象DTO*/
var mailService IMailService

// 用于对外暴露邮件服务，唯一的暴露点，供接口层调用
func GetMailService() IMailService {
	infras.Check(mailService)
	return mailService
}

// 服务具体实现初始化时设置服务对象，供核心业务层具体实现并设置
func SetMailService(service IMailService) {
	mailService = service
}

type IMailService interface {
	SendEmailForVerified(dto SendEmailForVerifiedDTO) error   // 绑定邮箱，发送验证邮件到指定邮箱
	SendEmailForgetPassword(SendEmailForgetPasswordDTO) error // 忘记密码，发送邮件到用户绑定的邮箱
}

type SendEmailForVerifiedDTO struct {
	ID    uint   `validate:"required,numeric"`
	Email string `validate:"required,email"`
}

type SendEmailForgetPasswordDTO struct {
	ID    uint   `validate:"required,numeric"`
	Email string `validate:"required,email"`
}
