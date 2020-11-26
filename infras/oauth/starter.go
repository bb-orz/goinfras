package oauth

import (
	"GoWebScaffold/infras"
)

var oauthManager *oAuthManager

type oAuthManager struct {
	Wechat *WechatOAuthManager
	Weibo  *WeiboOAuthManager
	QQ     *QQOAuthManager
}

func Manager() *oAuthManager {
	infras.Check(oauthManager)
	return oauthManager
}

type Starter struct {
	infras.BaseStarter
	cfg *Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("OAuth", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	oauthManager = new(oAuthManager)
	if s.cfg.QQSignSwitch {
		oauthManager.QQ = NewQQOauthManager(s.cfg)
		sctx.Logger().Info("QQ OAuth Manager Setup Successful!")
	}
	if s.cfg.WechatSignSwitch {
		oauthManager.Wechat = NewWechatOAuthManager(s.cfg)
		sctx.Logger().Info("Wechat OAuth Manager Setup Successful!")
	}
	if s.cfg.WeiboSignSwitch {
		oauthManager.Weibo = NewWeiboOAuthManager(s.cfg)
		sctx.Logger().Info("Weibo OAuth Manager Setup Successful!")
	}
}

func RunForTesting(config *Config) error {
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
