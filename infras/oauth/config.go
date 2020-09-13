package oauth

type OAuthConfig struct {
	WechatSignSwitch      bool   // 微信三方登录开关
	WechatOAuth2AppKey    string // 微信开发者appkey
	WechatOAuth2AppSecret string // 微信开发者secret
	WeiboSignSwitch       bool   // 微博三方登录开关
	WeiboOAuth2AppKey     string // 微博开发者appkey
	WeiboOAuth2AppSecret  string // 微博开发者secret
	QQSignSwitch          bool   // qq三方登录开关
	QQOAuth2AppKey        string // qq开发者appkey
	QQOAuth2AppSecret     string // qq开发者secret
}
