package oauth2


// OAuth2服务的通用鉴权接口
type OAuth2 interface{
	Authorize(code string) OAuthResult	// 根据微信返回的accessTokenCode开始鉴权并获取用户信息
	getAccessToken(code string) map[string]interface{}  // 从平台回调拿到accessTokenCode后调用此方法可拿到accessToken
	getUserInfo(accessToken string, openId string) map[string]interface{} // 拿到accessToken和openId后获取用户信息
}

// OAuth鉴权结果
type OAuthResult struct{
	Result bool
	UserInfo *OAuthUserInfo
}

// OAuth授权获取的用户信息
type OAuthUserInfo struct {
	AccessToken string
	OpenId string
	UnionId string
	NickName string
	Gender int
	AvatarUrl string
}

var QQOAuth2Manager QQOAuth
var WechatOAuth2Manager WechatOAuth
var WeiboOAuth2Manager WeiboOAuth

// 系统启动时初始化OAuth2服务
func InitOAuth2Manager()  {
	// 启用qq三方登录
	if config.QQSignSwitch {
		QQOAuth2Manager = QQOAuth{
			appKey:config.QQOAuth2AppKey,
			appSecret:config.QQOAuth2AppSecret,
		}
	}

	// 启用微信三方登录
	if config.WechatSignSwitch {
		WechatOAuth2Manager = WechatOAuth{
			appKey:config.WechatOAuth2AppKey,
			appSecret:config.WechatOAuth2AppSecret,
		}
	}

	// 启用微博三方登录
	if config.WeiboSignSwitch {
		WeiboOAuth2Manager = WeiboOAuth{
			appKey:config.WeiboOAuth2AppKey,
			appSecret:config.WeiboOAuth2AppSecret,
		}
	}

}
