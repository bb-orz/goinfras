package XOAuth

var wechatOM *WechatOAuthManager
var weiboOM *WeiboOAuthManager
var qqOM *QQOAuthManager

// 创建一个默认配置的Manager
func CreateDefaultManager(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	if config.QQSignSwitch {
		qqOM = NewQQOauthManager(config)
	}
	if config.WechatSignSwitch {
		wechatOM = NewWechatOAuthManager(config)
	}
	if config.WeiboSignSwitch {
		weiboOM = NewWeiboOAuthManager(config)
	}
}

func XQQOAuthManager() *QQOAuthManager {
	return qqOM
}

func XWeiboOAuthManager() *WeiboOAuthManager {
	return weiboOM
}

func XWechatOAuthManager() *WechatOAuthManager {
	return wechatOM
}
