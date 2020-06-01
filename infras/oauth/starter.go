package oauth

import (
	"GoWebScaffold/infras"
	"github.com/tietang/props/kvs"
)

var oauthManager *oAuthManager

type oAuthManager struct {
	Wechat *WechatOAuthManager
	Weibo  *WeiboOAuthManager
	QQ     *QQOAuthManager
}

func OAuthManager() *oAuthManager {
	infras.Check(oauthManager)
	return oauthManager
}

type OauthStarter struct {
	infras.BaseStarter
	cfg *oAuthConfig
}

func (s *OauthStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := oAuthConfig{}
	err := kvs.Unmarshal(configs, &define, "Oauth")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *OauthStarter) Setup(sctx *infras.StarterContext) {
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
