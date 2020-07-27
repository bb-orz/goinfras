package oauth2

import (
	"GoWebScaffold/core"
	"GoWebScaffold/infras/oauth"
)

/*
Oauth 领域层：实现第三方平台鉴权相关具体业务逻辑，主要为通过accessCode获取用户在第三方平台账号的信息
*/
type OauthDomain struct{}

func NewOauthDomain() *OauthDomain {
	domain := new(OauthDomain)
	return domain
}

// 通过accessCode获取qq user info
func (domain *OauthDomain) GetQQOauthUserInfo(accessCode string) (*oauth.OAuthAccountInfo, error) {
	result := oauth.OAuthManager().QQ.Authorize(accessCode)

	if result.Error != nil || !result.Result {
		return nil, core.WrapError(result.Error, core.ErrorFormatDomainThirdPart, "QQ.Authorize")
	}

	return result.UserInfo, nil
}

// 通过accessCode获取wechat user info
func (domain *OauthDomain) GetWechatOauthUserInfo(accessCode string) (*oauth.OAuthAccountInfo, error) {
	result := oauth.OAuthManager().Wechat.Authorize(accessCode)

	if result.Error != nil || !result.Result {
		return nil, core.WrapError(result.Error, core.ErrorFormatDomainThirdPart, "Wechat.Authorize")
	}

	return result.UserInfo, nil
}

// 通过accessCode获取weibo user info
func (domain *OauthDomain) GetWeiboOauthUserInfo(accessCode string) (*oauth.OAuthAccountInfo, error) {
	result := oauth.OAuthManager().Weibo.Authorize(accessCode)

	if result.Error != nil || !result.Result {
		return nil, core.WrapError(result.Error, core.ErrorFormatDomainThirdPart, "Weibo.Authorize")
	}

	return result.UserInfo, nil
}
