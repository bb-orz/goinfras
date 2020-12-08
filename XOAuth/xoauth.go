package XOAuth

import "goinfras"

var oAuthManager *OAuthManager

// 创建一个默认配置的Manager
func CreateDefaultManager(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	oAuthManager = new(OAuthManager)
	if config.QQSignSwitch {
		oAuthManager.QQOAuthManager = NewQQOauthManager(config)
	}
	if config.WechatSignSwitch {
		oAuthManager.WechatOAuthManager = NewWechatOAuthManager(config)
	}
	if config.WeiboSignSwitch {
		oAuthManager.WeiboOAuthManager = NewWeiboOAuthManager(config)
	}
}

type OAuthManager struct {
	WechatOAuthManager *WechatOAuthManager
	WeiboOAuthManager  *WeiboOAuthManager
	QQOAuthManager     *QQOAuthManager
}

func XManager() *OAuthManager {
	return oAuthManager
}

// 资源组件闭包执行
func XFManager(f func(m *OAuthManager) error) error {
	return f(oAuthManager)
}
