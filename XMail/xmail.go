package XMail

import (
	"gopkg.in/gomail.v2"
)

// 资源组件实例调用
func XDialer() *gomail.Dialer {
	return mailDialer
}

// 资源组件闭包执行
func XFDialer(f func(c *gomail.Dialer) error) error {
	return f(mailDialer)
}

// 邮件组件的通用操作实例
func XCommonMail() *CommonMail {
	c := new(CommonMail)
	c.dialer = XDialer()
	return c
}
