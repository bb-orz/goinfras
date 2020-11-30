package mail

import (
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/gomail.v2"
	"testing"
)

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
	m := NewNoAuthDialer(config.Server, config.Port)
	SetComponent(m)
}

func TestCommonMail(t *testing.T) {
	Convey("Test Common Mail", t, func() {
		TestingInstantiation(nil)

		// 组装邮件消息
		message := gomail.NewMessage(gomail.SetCharset("utf8"))
		message.SetAddressHeader("", "", "")
		message.SetBody("", "")

		// 邮件组件发送
		err := MailComponent().DialAndSend(message)
		So(err, ShouldBeNil)

	})
}
