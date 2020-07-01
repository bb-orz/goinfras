package mail

import (
	"GoWebScaffold/infras"
	"github.com/tietang/props/kvs"
	"gopkg.in/gomail.v2"
)

var mailDialer *gomail.Dialer

func MailDialer() *gomail.Dialer {
	infras.Check(mailDialer)
	return mailDialer
}

type MailStarter struct {
	infras.BaseStarter
	cfg *MailConfig
}

func (s *MailStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := MailConfig{}
	err := kvs.Unmarshal(configs, &define, "Mail")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *MailStarter) Setup(sctx *infras.StarterContext) {
	if s.cfg.NoAuth {
		mailDialer = NewNoAuthDialer(s.cfg.Server, s.cfg.Port)
	} else {
		mailDialer = NewAuthDialer(s.cfg.Server, s.cfg.User, s.cfg.Password, s.cfg.Port)
	}
}

/*For testing*/
func RunForTesting(config *MailConfig) error {
	var err error
	if config == nil {
		config = &MailConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err = p.Unmarshal(config)
		if err != nil {
			return err
		}
	}
	mailDialer = NewNoAuthDialer(config.Server, config.Port)
	return nil
}
