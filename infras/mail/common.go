package mail

import (
	"GoWebScaffold/infras/logger"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"io"
	"log"
	"time"
)

/*common 专门封装一些常用的操作*/

type CommonMail struct{}

func NewCommonMail() *CommonMail {
	return new(CommonMail)
}

// 本机发邮件
func (*CommonMail) SendMailNoSMTP(from, to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	s := gomail.SendFunc(func(from string, to []string, msg io.WriterTo) error {
		info := map[string]interface{}{
			"From":      from,
			"To":        to,
			"Subject":   subject,
			"ext/plain": body,
		}
		logger.CLogger().Info("Send Email:", zap.Any("Mail Info", info))
		return nil
	})

	if err := gomail.Send(s, m); err != nil {
		return err
	}

	return nil
}

// 发送简单邮件
func (*CommonMail) SendSimpleMail(from, to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	// m.SetAddressHeader()
	m.SetBody("text/plain", body)

	return MailComponent().DialAndSend(m)
}

// 使用通道在窗口时间内批量发送邮件
func (*CommonMail) SendBatchMails(msgCh <-chan *gomail.Message, duration time.Duration) error {
	go func() {
		var s gomail.SendCloser
		var err error
		open := false // 拨号状态
		for {
			select {
			case m, ok := <-msgCh:
				if !ok {
					return
				}
				if !open {
					if s, err = MailComponent().Dial(); err != nil {
						panic(err)
					}
					open = true
				}

				if err := gomail.Send(s, m); err != nil {
					log.Print(err)
				}
			case <-time.After(duration):
				// 超时之前保持连接
				if open {
					if err := s.Close(); err != nil {
						panic(err)
					}
					open = false
				}
			}
		}
	}()

	return nil
}
