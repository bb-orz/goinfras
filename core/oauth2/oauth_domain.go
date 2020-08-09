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
	var oAuthResult oauth.OAuthResult
	oAuthResult = oauth.OAuthManager().QQ.Authorize(accessCode)

	if oAuthResult.Error != nil || !oAuthResult.Result {
		return nil, core.WrapError(oAuthResult.Error, core.ErrorFormatDomainThirdPart, "QQ.Authorize")
	}

	return oAuthResult.UserInfo, nil
}

// 通过accessCode获取wechat user info
func (domain *OauthDomain) GetWechatOauthUserInfo(accessCode string) (*oauth.OAuthAccountInfo, error) {
	var oAuthResult oauth.OAuthResult
	oAuthResult = oauth.OAuthManager().Wechat.Authorize(accessCode)

	if oAuthResult.Error != nil || !oAuthResult.Result {
		return nil, core.WrapError(oAuthResult.Error, core.ErrorFormatDomainThirdPart, "Wechat.Authorize")
	}

	return oAuthResult.UserInfo, nil
}

// 通过accessCode获取weibo user info
func (domain *OauthDomain) GetWeiboOauthUserInfo(accessCode string) (*oauth.OAuthAccountInfo, error) {
	var oAuthResult oauth.OAuthResult
	oAuthResult = oauth.OAuthManager().Weibo.Authorize(accessCode)

	if oAuthResult.Error != nil || !oAuthResult.Result {
		return nil, core.WrapError(oAuthResult.Error, core.ErrorFormatDomainThirdPart, "Weibo.Authorize")
	}

	return oAuthResult.UserInfo, nil
}
