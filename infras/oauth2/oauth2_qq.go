package oauth2

import (
	"encoding/json"
	"github.com/imroc/req"
	"strings"
)

const qqGetAccessTokenUrl = "https://graph.qq.com/oauth2.0/token"
const qqOpenIdUrl = "https://graph.qq.com/oauth2.0/me"
const qqGetUserInfoUrl = "https://graph.qq.com/user/get_user_info"

// 实现QQOAuth2鉴权
type QQOAuth struct {
	appKey      string
	appSecret   string
	redirectUrl string
}

func (oauth *QQOAuth) Authorize(code string) OAuthResult {
	// 先获取accessToken
	accessTokenResp := oauth.getAccessToken(code)
	if accessTokenResp == nil {
		return OAuthResult{false, nil}
	}
	accessToken := accessTokenResp["access_token"].(string)

	// 再获取openId和unionId
	openidResp := oauth.getOpenId(accessToken)
	if _, ok := openidResp["error"]; ok {
		return OAuthResult{false, nil}
	}
	openId := openidResp["openid"].(string)
	unionId := openidResp["unionid"].(string)

	// 最后获取用户信息
	userInfoMap := oauth.getUserInfo(accessToken, openId)
	if userInfoMap == nil {
		return OAuthResult{false, nil}
	}
	var genderN int
	gender, ok := userInfoMap["gender"]
	if !ok {
		genderN = 1
	}
	if gender.(string) == "女" {
		genderN = 2
	} else {
		genderN = 1
	}
	return OAuthResult{true, &OAuthUserInfo{
		accessToken,
		openId,
		unionId,
		userInfoMap["nickname"].(string),
		genderN,
		userInfoMap["figureurl_qq_1"].(string),
	}}
}

func (oauth *QQOAuth) getAccessToken(code string) map[string]interface{} {

	params := req.Param{
		"grant_type":    "authorization_code",
		"client_id":     oauth.appKey,
		"client_secret": oauth.appSecret,
		"code":          code,
		"redirect_uri":  oauth.redirectUrl,
	}
	resp, err := req.Get(qqGetAccessTokenUrl, params)
	if err != nil {
		return nil
	}
	response := resp.String()

	if strings.Contains(response, "callback") {
		return nil
	}
	temp := strings.Split(response, "&")[0]
	accessToken := strings.Split(temp, "=")[1]
	return map[string]interface{}{"access_token": accessToken}
}

func (oauth *QQOAuth) getOpenId(accessToken string) map[string]interface{} {
	params := req.Param{
		"access_token": accessToken,
		"unionid":      1,
	}

	resp, err := req.Get(qqOpenIdUrl, params)
	if err != nil {
		return nil
	}

	respStr := resp.String()
	var response map[string]interface{}
	err = json.Unmarshal([]byte(respStr[10:len(respStr)-3]), &response)
	if err != nil {
		return nil
	}

	return response
}

/*
// 返回用户信息结果示例

{
"ret":0,
"msg":"",
"nickname":"Peter",
"figureurl":"http://qzapp.qlogo.cn/qzapp/111111/942FEA70050EEAFBD4DCE2C1FC775E56/30",
"figureurl_1":"http://qzapp.qlogo.cn/qzapp/111111/942FEA70050EEAFBD4DCE2C1FC775E56/50",
"figureurl_2":"http://qzapp.qlogo.cn/qzapp/111111/942FEA70050EEAFBD4DCE2C1FC775E56/100",
"figureurl_qq_1":"http://q.qlogo.cn/qqapp/100312990/DE1931D5330620DBD07FB4A5422917B6/40",
"figureurl_qq_2":"http://q.qlogo.cn/qqapp/100312990/DE1931D5330620DBD07FB4A5422917B6/100",
"gender":"男",
"is_yellow_vip":"1",
"vip":"1",
"yellow_vip_level":"7",
"level":"7",
"is_yellow_year_vip":"1"
}
*/
func (oauth *QQOAuth) getUserInfo(accessToken string, openId string) map[string]interface{} {
	params := req.Param{
		"access_token":       accessToken,
		"oauth_consumer_key": oauth.appKey,
		"openid":             openId,
	}

	resp, err := req.Get(qqGetUserInfoUrl, params)
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
