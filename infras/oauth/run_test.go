package oauth

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

	oauthManager = new(oAuthManager)
	if config.QQSignSwitch {
		oauthManager.QQ = NewQQOauthManager(config)
	}
	if config.WechatSignSwitch {
		oauthManager.Wechat = NewWechatOAuthManager(config)
	}
	if config.WeiboSignSwitch {
		oauthManager.Weibo = NewWeiboOAuthManager(config)
	}

	return nil
}
