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
	QQLogin(accessCode string)     // qq三方登录
	WeixinLogin(accessCode string) // 微信三方登录
	WeiboLogin(accessCode string)  // 微博三方登录
}
