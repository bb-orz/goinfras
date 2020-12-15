package XMail

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"io"
	"testing"
)

// 非本地服务器，通过API发送邮件
func TestCommonSendNoSMTPMail(t *testing.T) {
	Convey("Test Common Mail", t, func() {
		CreateDefaultManager(nil)

		err := XCommonMail().SendMailNoSMTP("goinras@xxx.com", "test", "test XMail", BodyTypePlain, []string{"gofuncchan@163.com"}, func(from string, to []string, msg io.WriterTo) error {
			fmt.Println("From:", from)
			fmt.Println("To:", to)
			fmt.Println("Msg:", msg)

			// TODO 通过外部API发送邮件

			return nil
		})

		So(err, ShouldBeNil)

	})
}

// 发送简单邮件，测试前请先设置默认配置信息
func TestCommonSendSimpleMail(t *testing.T) {
	Convey("Test Common Mail", t, func() {
		CreateDefaultManager(nil)

		// 发送
		err := XCommonMail().SendSimpleMail(
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			[]string{""},
		)
		So(err, ShouldBeNil)

	})
}

// 群发邮件，测试前请先设置默认配置信息
func TestCommonSendNewsLetter(t *testing.T) {
	Convey("Test Common Mail", t, func() {
		CreateDefaultManager(nil)

		receivers := []NewsLetterReceiver{
			{
				Name:    "",
				Address: "",
			},
			{
				Name:    "",
				Address: "",
			},
		}

		err := XCommonMail().SendNewsLetter(receivers, "infras@xxx.com", "test new letter email", "test new letter email", BodyTypePlain)
		So(err, ShouldBeNil)
	})
}

//
func TestCommonSendBatchMails(t *testing.T) {
	Convey("Test Common Mail", t, func() {
		CreateDefaultManager(nil)
		msgCh := make(chan *gomail.Message)
		defer func() {
			close(msgCh)
		}()
		err := XCommonMail().SendBatchMails(msgCh, 10)
		So(err, ShouldBeNil)

		// TODO Send Message to msgCh
		msg1 := gomail.NewMessage()
		msg1.SetHeader("From", "")
		msg1.SetAddressHeader("To", "", "")
		msg1.SetHeader("Subject", "Newsletter #1")
		msg1.SetBody("text/plain", "")
		msgCh <- msg1

		// continue...
	})
}

func TestStarter(t *testing.T) {
	Convey("Test XMail Starter", t, func() {
		s := NewStarter()
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s.Init(sctx)
		Println("Starter Init Successful!")
		s.Setup(sctx)
		Println("Starter Setup Successful!")

		if s.Check(sctx) {
			Println("Component Check Successful!")
		} else {
			Println("Component Check Fail!")
		}

		err = XCommonMail().SendSimpleMail(
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			[]string{""},
		)

		So(err, ShouldBeNil)

	})
}
