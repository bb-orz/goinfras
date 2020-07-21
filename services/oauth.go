package services

import "GoWebScaffold/infras"

/* 定义三方登录模块的服务层方法，并定义数据传输对象DTO*/

var oauthService IOAuthService

func GetOAuthService() IOAuthService {
	infras.Check(oauthService)
	return oauthService
}

func SetOAuthService(service IOAuthService) {
	oauthService = service
}

type IOAuthService interface {
	QQOAuth(dto QQLoginDTO) (string, error)         // qq三方登录
	WeixinOAuth(dto WeixinLoginDTO) (string, error) // 微信三方登录
	WeiboOAuth(dto WeiboLoginDTO) (string, error)   // 微博三方登录
}

type QQLoginDTO struct {
	accessCode string
}

type WeixinLoginDTO struct {
	accessCode string
}

type WeiboLoginDTO struct {
	accessCode string
}
