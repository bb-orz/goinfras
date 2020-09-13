package mail

import (
	"GoWebScaffold/infras"
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
	viper := sctx.Configs()
	define := MailConfig{}
	err := viper.UnmarshalKey("Mail", &define)
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
func RunForTesting(config *MailConfig) {
	if config == nil {
		config = &MailConfig{
			NoAuth:   false,                   // 使用本地SMTP服务器发送电子邮件。
			NoSmtp:   false,                   // 使用API​​或后缀发送电子邮件。
			Server:   "smtp.qq.com",           // 使用外部SMTP服务器
			Port:     587,                     // 外部SMTP服务端口
			User:     "your qq mail account",  // 你的三方邮箱地址
			Password: "your qq mail password", // 你的邮箱密码
		}

	}
	// mailDialer = NewNoAuthDialer(config.Server, config.Port)
}
