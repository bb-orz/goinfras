package user

import (
	"GoWebScaffold/core"
	"GoWebScaffold/infras/oauth"
	"GoWebScaffold/services"
)

/*
Oauth 领域层：实现三方登录相关具体业务逻辑
封装领域层的错误信息并返回给调用者
*/
type OauthDomain struct {
	dao   *oauthDAO
	cache *oauthCache
}

func NewOauthDomain() *OauthDomain {
	domain := new(OauthDomain)
	domain.dao = NewOauthDAO()
	domain.cache = NewOauthCache()
	return domain
}

// 通过accessCode获取qq user info
func (domain *OauthDomain) GetQQOauthUserInfo(accessCode string) (*oauth.OAuthUserInfo, error) {
	result := oauth.OAuthManager().QQ.Authorize(accessCode)

	if result.Error != nil || !result.Result {
		return nil, core.WrapError(result.Error, core.ErrorFormatDomainThirdPart, "QQ.Authorize")
	}

	return result.UserInfo, nil
}

// 通过accessCode获取wechat user info
func (domain *OauthDomain) GetWechatOauthUserInfo(accessCode string) (*oauth.OAuthUserInfo, error) {
	result := oauth.OAuthManager().Wechat.Authorize(accessCode)

	if result.Error != nil || !result.Result {
		return nil, core.WrapError(result.Error, core.ErrorFormatDomainThirdPart, "Wechat.Authorize")
	}

	return result.UserInfo, nil
}

// 通过accessCode获取weibo user info
func (domain *OauthDomain) GetWeiboOauthUserInfo(accessCode string) (*oauth.OAuthUserInfo, error) {
	result := oauth.OAuthManager().Weibo.Authorize(accessCode)

	if result.Error != nil || !result.Result {
		return nil, core.WrapError(result.Error, core.ErrorFormatDomainThirdPart, "Weibo.Authorize")
	}

	return result.UserInfo, nil
}

// 查找Oauth三方注册账号是否存在
func (domain *OauthDomain) GetOauthUserBinding(platform uint, openId, unionId string) (*services.OauthInfoDTO, error) {

	return nil, nil
}

// 创建Oauth三方账号绑定
func (domain *OauthDomain) CreateOauthUserBinding(platform uint) {
	userOauthModle := UserOauth{}

}
