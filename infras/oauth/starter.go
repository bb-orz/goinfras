package oauth

import (
	"GoWebScaffold/infras"
)

type Starter struct {
	infras.BaseStarter
	cfg Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("OAuth", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var om *OAuthManager
	om = new(OAuthManager)
	if s.cfg.QQSignSwitch {
		om.QQ = NewQQOauthManager(&s.cfg)
		sctx.Logger().Info("QQ OAuth Manager Setup Successful!")
	}
	if s.cfg.WechatSignSwitch {
		om.Wechat = NewWechatOAuthManager(&s.cfg)
		sctx.Logger().Info("Wechat OAuth Manager Setup Successful!")
	}
	if s.cfg.WeiboSignSwitch {
		om.Weibo = NewWeiboOAuthManager(&s.cfg)
		sctx.Logger().Info("Weibo OAuth Manager Setup Successful!")
	}

	SetComponent(om)
}
