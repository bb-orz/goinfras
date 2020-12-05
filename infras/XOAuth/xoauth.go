package XOAuth

var oAuthManager *OAuthManager

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

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	if config == nil {
		config = &Config{
			false,
			"",
			"",
			false,
			"",
			"",
			false,
			"",
			"",
		}
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

	return nil
}
