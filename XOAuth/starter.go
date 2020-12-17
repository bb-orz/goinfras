package XOAuth

import (
	"fmt"
	"github.com/bb-orz/goinfras"
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
	fmt.Printf("Print XOAuth Config: %v", *define)
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	if s.cfg.QQSignSwitch {
		qqOM = NewQQOauthManager(s.cfg)
	}
	if s.cfg.WechatSignSwitch {
		wechatOM = NewWechatOAuthManager(s.cfg)
	}
	if s.cfg.WeiboSignSwitch {
		weiboOM = NewWeiboOAuthManager(s.cfg)
	}

}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	var err error
	if s.cfg.QQSignSwitch {
		if err = goinfras.Check(qqOM); err != nil {
			sctx.Logger().Error(fmt.Sprintf("[%s Starter]: QQ OAuth Manager Setup Fail!", s.Name()))
			return false
		}
		sctx.Logger().Info(fmt.Sprintf("[%s Starter]:QQ OAuth Manager Setup Successful!", s.Name()))
	}

	if s.cfg.WechatSignSwitch {
		if err = goinfras.Check(wechatOM); err != nil {
			sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Wechat OAuth Manager Setup Fail!", s.Name()))
			return false
		}
		sctx.Logger().Info(fmt.Sprintf("[%s Starter]:Wechat OAuth Manager Setup Successful!", s.Name()))
	}

	if s.cfg.WeiboSignSwitch {
		if err = goinfras.Check(weiboOM); err != nil {
			sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Weibo OAuth Manager Setup Fail!", s.Name()))
			return false
		}
		sctx.Logger().Info(fmt.Sprintf("[%s Starter]:Weibo OAuth Manager Setup Successful!", s.Name()))
	}

	return true
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.AppGroup }
