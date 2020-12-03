package XMail

import (
	"gopkg.in/gomail.v2"
)

var mailDialer *gomail.Dialer

// 资源组件实例调用
func XDialer() *gomail.Dialer {
	return mailDialer
}

// 资源组件闭包执行
func XFDialer(f func(c *gomail.Dialer) error) error {
	return f(mailDialer)
}
