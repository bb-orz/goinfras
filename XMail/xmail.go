package XMail

import (
	"gopkg.in/gomail.v2"
)

var mailDialer *gomail.Dialer

// 创建一个默认配置的Manager
func CreateDefaultManager(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	mailDialer = NewNoAuthDialer(config.Server, config.Port)
}

// 创建一个默认配置的Manager
func CreateDefaultAuthManager(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	mailDialer = NewAuthDialer(config.Server, config.User, config.Password, config.Port)
}

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