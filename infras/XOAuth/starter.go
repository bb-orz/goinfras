package XOAuth

import (
	"GoWebScaffold/infras"
	"fmt"
)

type starter struct {
	infras.BaseStarter
	cfg Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = Config{}
	return starter
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
		oAuthManager.QQOAuthManager = NewQQOauthManager(&s.cfg)
	}
	if s.cfg.WechatSignSwitch {
		oAuthManager.WechatOAuthManager = NewWechatOAuthManager(&s.cfg)
	}
	if s.cfg.WeiboSignSwitch {
		oAuthManager.WeiboOAuthManager = NewWeiboOAuthManager(&s.cfg)
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