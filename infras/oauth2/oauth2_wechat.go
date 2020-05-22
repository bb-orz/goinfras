package oauth2

import (
	"github.com/imroc/req"
)

const wechatGetAccessTokenUrl = "https://api.weixin.qq.com/sns/oauth2/access_token"
const wechatGetUserInfoUrl = "https://api.weixin.qq.com/sns/userinfo"

// 实现 微信 OAuth2鉴权
type WechatOAuth struct {
	appKey      string
	appSecret   string
	redirectUrl string
}

func (oauth *WechatOAuth) Authorize(code string) OAuthResult {
	accessTokenResp := oauth.getAccessToken(code)
	if accessTokenResp == nil {
		return OAuthResult{false, nil}
	}

	// 获取accessToken接口返回错误码
	_, ok := accessTokenResp["errcode"]
	if ok {
		return OAuthResult{false, nil}
	}

	openId := accessTokenResp["openid"].(string)
	accessToken := accessTokenResp["access_token"].(string)
	userInfoMap := oauth.getUserInfo(accessToken, openId)
	if userInfoMap == nil {
		return OAuthResult{false, nil}
	}

	// 获取用户信息接口返回错误码
	_, ok = userInfoMap["errcode"]
	if ok {
		return OAuthResult{false, nil}
	}

	return OAuthResult{true, &OAuthUserInfo{
		accessToken,
		openId,
		userInfoMap["unionid"].(string),
		userInfoMap["nickname"].(string),
		userInfoMap["sex"].(int),
		userInfoMap["figureurl_qq_1"].(string),
	}}
}

func (oauth *WechatOAuth) getAccessToken(code string) map[string]interface{} {
	params := req.Param{
		"appid":      oauth.appKey,
		"secret":     oauth.appSecret,
		"code":       code,
		"grant_type": "authorization_code",
	}
	resp, err := req.Get(wechatGetAccessTokenUrl, params)
	if err != nil {
		return nil
	}
	var response map[string]interface{}
	err = resp.ToJSON(&response)
	if err != nil {
		return nil
	}
	return response

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
func (oauth *WechatOAuth) getUserInfo(accessToken string, openId string) map[string]interface{} {
	params := req.Param{
		"access_token": accessToken,
		"openid":       openId,
	}
	resp, err := req.Get(wechatGetUserInfoUrl, params)
	if err != nil {
		return nil
	}
	var response map[string]interface{}
	err = resp.ToJSON(&response)
	if err != nil {
		return nil
	}
	return response

}
