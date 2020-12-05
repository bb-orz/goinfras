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

// 邮件组件的通用操作实例
func XCommonMail() *CommonMail {
	c := new(CommonMail)
	c.dialer = XDialer()
	return c
}

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) {
	if config == nil {
		config = &Config{
			NoAuth:   false,                   // 使用本地SMTP服务器发送电子邮件。
			NoSmtp:   false,                   // 使用API​​或后缀发送电子邮件。
			Server:   "smtp.qq.com",           // 使用外部SMTP服务器
			Port:     587,                     // 外部SMTP服务端口
			User:     "your qq mail account",  // 你的三方邮箱地址
			Password: "your qq mail password", // 你的邮箱密码
		}

	}
	mailDialer = NewNoAuthDialer(config.Server, config.Port)

}
