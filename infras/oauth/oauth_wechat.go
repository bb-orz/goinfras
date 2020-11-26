package oauth

import (
	"errors"
	"github.com/imroc/req"
)

const wechatGetAccessTokenUrl = "https://api.weixin.qq.com/sns/oauth2/access_token"
const wechatGetUserInfoUrl = "https://api.weixin.qq.com/sns/userinfo"

// 实现 微信 OAuth2鉴权
type WechatOAuthManager struct {
	appKey      string
	appSecret   string
	redirectUrl string
}

func NewWechatOAuthManager(cfg *Config) *WechatOAuthManager {
	return &WechatOAuthManager{
		appKey:    cfg.WechatOAuth2AppKey,
		appSecret: cfg.WechatOAuth2AppSecret,
	}
}

func (oauth *WechatOAuthManager) Authorize(code string) OAuthResult {
	accessTokenResp, err := oauth.getAccessToken(code)
	if err != nil || accessTokenResp == nil {
		return OAuthResult{false, nil, err}
	}

	// 获取accessToken接口返回错误码
	e, ok := accessTokenResp["errcode"]
	if ok {
		return OAuthResult{false, nil, errors.New(e.(string))}
	}

	openId := accessTokenResp["openid"].(string)
	accessToken := accessTokenResp["access_token"].(string)
	userInfoMap, err := oauth.getUserInfo(accessToken, openId)
	if err != nil || userInfoMap == nil {
		return OAuthResult{false, nil, err}
	}

	// 获取用户信息接口返回错误码
	e, ok = userInfoMap["errcode"]
	if ok {
		return OAuthResult{false, nil, errors.New(e.(string))}
	}

	return OAuthResult{true, &OAuthAccountInfo{
		accessToken,
		openId,
		userInfoMap["unionid"].(string),
		userInfoMap["nickname"].(string),
		userInfoMap["sex"].(uint),
		userInfoMap["figureurl_qq_1"].(string),
	}, nil}
}

func (oauth *WechatOAuthManager) getAccessToken(code string) (map[string]interface{}, error) {
	params := req.Param{
		"appid":      oauth.appKey,
		"secret":     oauth.appSecret,
		"code":       code,
		"grant_type": "authorization_code",
	}
	resp, err := req.Get(wechatGetAccessTokenUrl, params)
	if err != nil {
		return nil, err
	}
	var response map[string]interface{}
	err = resp.ToJSON(&response)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
// 返回用户信息结果示例
{
"openid":"OPENID",
"nickname":"NICKNAME",
"sex":1,
"province":"PROVINCE",
"city":"CITY",
"country":"COUNTRY",
"headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0",
"privilege":[
"PRIVILEGE1",
"PRIVILEGE2"
],
"unionid": " o6_bmasdasdsad6_2sgVt7hMZOPfL"

}
*/
func (oauth *WechatOAuthManager) getUserInfo(accessToken string, openId string) (map[string]interface{}, error) {
	params := req.Param{
		"access_token": accessToken,
		"openid":       openId,
	}
	resp, err := req.Get(wechatGetUserInfoUrl, params)
	if err != nil {
		return nil, err
	}
	var response map[string]interface{}
	err = resp.ToJSON(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
