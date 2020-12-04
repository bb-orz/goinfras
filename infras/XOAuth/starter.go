package XOAuth

import (
	"GoWebScaffold/infras"
	"fmt"
)

type starter struct {
	infras.BaseStarter
	cfg Config
}

func (s *starter) Name() string {
	return "XOAuth"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("OAuth", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	oAuthManager = new(OAuthManager)
	if s.cfg.QQSignSwitch {
		oAuthManager.QQ = NewQQOauthManager(&s.cfg)
		sctx.Logger().Info("QQ OAuth Manager Setup Successful!")
	}
	if s.cfg.WechatSignSwitch {
		oAuthManager.Wechat = NewWechatOAuthManager(&s.cfg)
		sctx.Logger().Info("Wechat OAuth Manager Setup Successful!")
	}
	if s.cfg.WeiboSignSwitch {
		oAuthManager.Weibo = NewWeiboOAuthManager(&s.cfg)
		sctx.Logger().Info("Weibo OAuth Manager Setup Successful!")
	}

}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(oAuthManager)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: OAuth Manager Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: OAuth Manager Setup Successful!", s.Name()))
	return true
}
