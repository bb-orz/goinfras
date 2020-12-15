package XMail

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"go.uber.org/zap"
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
	return "XMail"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Cron", &define)
		goinfras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print Mail Config:", zap.Any("MailConfig", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	if s.cfg.NoAuth {
		mailDialer = NewNoAuthDialer(s.cfg.Server, s.cfg.Port)
	} else {
		mailDialer = NewAuthDialer(s.cfg.Server, s.cfg.User, s.cfg.Password, s.cfg.Port)
	}
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(mailDialer)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Mail Dialer Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Mail Dialer Setup Successful!", s.Name()))
	return true
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.ResourcesGroup }
