package XMail

import (
	"gopkg.in/gomail.v2"
	"log"
	"time"
)

const (
	BodyTypePlain = "text/plain" // 纯文本
	BodyTypeHTML  = "text/html"  // HTML格式
)

/*common 专门封装一些常用的操作*/

type CommonMail struct {
	dialer *gomail.Dialer
}

/**
 * @Description: 非本地邮件服务器，通过API发送邮件
 * @receiver c
 * @param from 发送方
 * @param subject 邮件主题
 * @param body 邮件主体
 * @param bodyType 邮件主体格式：BodyTypePlain(文本格式)或BodyTypeHTML(HTML格式)
 * @param to 接收方
 * @param sendFunc 发送处理函数,闭包执行非本服务器的API发送邮件
 * @return error
 */
func (c *CommonMail) SendMailNoSMTP(from, subject, body, bodyType string, to []string, sendFunc gomail.SendFunc) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", subject)
	msg.SetBody(bodyType, body)

	err := gomail.Send(sendFunc, msg)
	return err
}

/**
 * @Description: 使用SMTP服务器发送简单邮件
 * @receiver c
 * @param from 发送方
 * @param to 接收方（可多个）
 * @param ccAddress 抄送地址
 * @param ccName 抄送名称
 * @param subject 邮件主题
 * @param body  邮件主体
 * @param bodyType 邮件主体格式：BodyTypePlain(文本格式)或BodyTypeHTML(HTML格式)
 * @param attach 附件文件资源地址
 * @return error
 */
func (c *CommonMail) SendSimpleMail(from, ccAddress, ccName, subject, body, bodyType, attach string, to []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetAddressHeader("Cc", ccAddress, ccName)
	m.SetBody(bodyType, body)

	return c.dialer.DialAndSend(m)
}

type NewsLetterReceiver struct {
	Name    string
	Address string
}

/**
 * @Description: 使用SMTP服务器批量发送邮件
 * @receiver c
 * @param receivers 接收者
 * @param from 发送方
 * @param subject 邮件主题
 * @param body 邮件主体
 * @param bodyType 邮件主体格式：BodyTypePlain(文本格式)或BodyTypeHTML(HTML格式)
 * @return error
 */
func (c *CommonMail) SendNewsLetter(receivers []NewsLetterReceiver, from, subject, body, bodyType string) error {
	msg := gomail.NewMessage()
	for _, r := range receivers {
		msg.SetHeader("From", from)
		msg.SetAddressHeader("To", r.Address, r.Name)
		msg.SetHeader("Subject", "Newsletter #1")
		msg.SetBody(bodyType, body)

		sendCloser, err := XDialer().Dial()
		if err != nil {
			return err
		}
		if err := gomail.Send(sendCloser, msg); err != nil {
			log.Printf("Could not send email to %q: %v", r.Address, err)
		}
		msg.Reset()
	}
	return nil
}

/**
 * @Description: 使用SMTP服务器，使用通道在窗口时间内批量发送邮件
 * @receiver c
 * @param msgCh 发送邮件信息的通道
 * @param duration 发送超时时间
 * @return error
 */
func (c *CommonMail) SendBatchMails(msgCh <-chan *gomail.Message, duration time.Duration) error {
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
					if s, err = c.dialer.Dial(); err != nil {
						panic(err)
					}
					open = true
				}

				if err := gomail.Send(s, m); err != nil {
					log.Print(err)
				}
			case <-time.After(duration):
				// 超时之后关闭发送连接
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
