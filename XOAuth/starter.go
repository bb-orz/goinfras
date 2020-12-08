package XOAuth

import (
	"fmt"
	"go.uber.org/zap"
	"goinfras"
)

type starter struct {
	goinfras.BaseStarter
	cfg *Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	return starter
}

func (s *starter) Name() string {
	return "XOAuth"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("OAuth", &define)
		goinfras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print OAuth Config:", zap.Any("OAuthConfig", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	oAuthManager = new(OAuthManager)
	if s.cfg.QQSignSwitch {
		oAuthManager.QQOAuthManager = NewQQOauthManager(s.cfg)
	}
	if s.cfg.WechatSignSwitch {
		oAuthManager.WechatOAuthManager = NewWechatOAuthManager(s.cfg)
	}
	if s.cfg.WeiboSignSwitch {
		oAuthManager.WeiboOAuthManager = NewWeiboOAuthManager(s.cfg)
	}

}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(oAuthManager)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: OAuth Manager Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: OAuth Manager Setup Successful!", s.Name()))
	return true
}
