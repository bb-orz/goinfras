package core

import (
	"GoWebScaffold/core/user"
	"GoWebScaffold/infras/validate"
	"GoWebScaffold/services"
	"fmt"
	"sync"
)

var _ services.IOAuthService = new(OauthService)

func init() {
	// 初始化该业务模块时实例化服务
	var once sync.Once
	once.Do(func() {
		oauthService := new(OauthService)
		oauthService.userDomain = user.NewUserDomain()
		oauthService.oauthDomain = user.NewOauthDomain()
		services.SetOAuthService(oauthService)
	})
}

type OauthService struct {
	userDomain  *user.UserDomain
	oauthDomain *user.OauthDomain
}

func (service *OauthService) QQOAuth(dto services.QQLoginDTO) (string, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return "", WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// oauth domain：使用qq回调授权码code开始鉴权流程并获取QQ用户信息
	authInfo, err := service.oauthDomain.GetQQUserInfo(dto.AccessCode)
	if err != nil {
		return "", WrapError(err, ErrorFormatServiceNetRequest, "GetQQUserInfo")
	}

	fmt.Printf(authInfo.NickName)

	// oauth domain: 使用OpenId UnionId查找user oauth表查看用户是否存在
	oauthInfo, err := service.oauthDomain.GetOauthUser(authInfo.OpenId, authInfo.UnionId)
	if err != nil {
		return "", WrapError(err, ErrorFormatServiceStorage, "IsOauthUserExist")
	}

	if oauthInfo != nil {
		// 如存在进入登录流程
	} else {
		// 如不存在进入创建用户流程

	}

	// 返回jwt

	return "", nil
}

func (service *OauthService) WechatOAuth(dto services.WechatLoginDTO) (string, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return "", WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// oauth domain：使用qq回调授权码code开始鉴权流程并获取QQ用户信息

	// oauth domain: 使用OpenId UnionId查找user oauth表查看用户是否存在

	// 不存在进入创建用户流程，存在进入登录流程

	// 返回jwt

	return "", nil
}

func (service *OauthService) WeiboOAuth(dto services.WeiboLoginDTO) (string, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return "", WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// oauth domain：使用qq回调授权码code开始鉴权流程并获取QQ用户信息

	// oauth domain: 使用OpenId UnionId查找user oauth表查看用户是否存在

	// 不存在进入创建用户流程，存在进入登录流程

	// 返回jwt

	return "", nil
}
