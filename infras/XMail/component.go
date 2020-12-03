package XMail

import (
	"GoWebScaffold/infras"
	"gopkg.in/gomail.v2"
)

var mailDialer *gomail.Dialer

func MailComponent() *gomail.Dialer {
	infras.Check(mailDialer)
	return mailDialer
}

func SetComponent(m *gomail.Dialer) {
	mailDialer = m
}
