package XMail

import (
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/gomail.v2"
	"testing"
)

func TestCommonMail(t *testing.T) {
	Convey("Test Common Mail", t, func() {
		TestingInstantiation(nil)

		// 组装邮件消息
		message := gomail.NewMessage(gomail.SetCharset("utf8"))
		message.SetAddressHeader("", "", "")
		message.SetBody("", "")

		// 邮件组件发送
		err := XDialer().DialAndSend(message)
		So(err, ShouldBeNil)

	})
}
