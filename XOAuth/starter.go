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
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().SDebug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %v \n", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	if s.cfg.QQSignSwitch {
		qqOM = NewQQOauthManager(s.cfg)
		sctx.Logger().SDebug(s.Name(), goinfras.StepSetup, fmt.Sprintf("QQ OAuth Manager Steuped!  \n"))
	}
	if s.cfg.WechatSignSwitch {
		wechatOM = NewWechatOAuthManager(s.cfg)
		sctx.Logger().SDebug(s.Name(), goinfras.StepSetup, fmt.Sprintf("Wechat OAuth Manager Steuped!  \n"))
	}
	if s.cfg.WeiboSignSwitch {
		weiboOM = NewWeiboOAuthManager(s.cfg)
		sctx.Logger().SDebug(s.Name(), goinfras.StepSetup, fmt.Sprintf("Weibo OAuth Manager Steuped!  \n"))
	}

}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	var err error
	if s.cfg.QQSignSwitch {
		err = goinfras.Check(qqOM)
		if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
			sctx.Logger().SInfo(s.Name(), goinfras.StepCheck, fmt.Sprintf("QQ OAuth Manager Steup Successful! \n"))
		}
	}

	if s.cfg.WechatSignSwitch {
		err = goinfras.Check(wechatOM)
		if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
			sctx.Logger().SInfo(s.Name(), goinfras.StepCheck, fmt.Sprintf("Wechat OAuth Manager Steup Successful! \n"))
		}
	}

	if s.cfg.WeiboSignSwitch {
		err = goinfras.Check(weiboOM)
		if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
			sctx.Logger().SInfo(s.Name(), goinfras.StepCheck, fmt.Sprintf("Weibo OAuth Manager Steup Successful! \n"))
		}
	}

	return true
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.AppGroup }
