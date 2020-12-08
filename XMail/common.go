package XMail

import (
	"bytes"
	"go.uber.org/zap"
	"goinfras/XLogger"
	"gopkg.in/gomail.v2"
	"io"
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
 * @Description: NoSMTP发邮件
 * @receiver c
 * @param from 发送方
 * @param subject 邮件主题
 * @param body 邮件主体
 * @param bodyType 邮件主体格式：BodyTypePlain(文本格式)或BodyTypeHTML(HTML格式)
 * @param to 接收方
 * @param sendFunc 发送处理函数
 * @return error
 */
func (c *CommonMail) SendMailNoSMTP(from, subject, body, bodyType string, to []string, sendFunc SendFunc) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", subject)
	msg.SetBody(bodyType, body)

	if sendFunc == nil {
		sendFunc = defaultSendFunc
	}
	sf := gomail.SendFunc(sendFunc)

	if err := gomail.Send(sf, msg); err != nil {
		return err
	}

	return nil
}

type SendFunc func(from string, to []string, msg io.WriterTo) error

var defaultSendFunc = func(from string, to []string, msg io.WriterTo) error {
	var msgBuf []byte
	var err error
	buf := bytes.NewBuffer(msgBuf)
	_, err = msg.WriteTo(buf)
	if err != nil {
		return err
	}

	mailInfo := map[string]interface{}{
		"From": from,
		"To":   to,
		"msg":  msgBuf,
	}

	XLogger.XCommon().Info("Send NoSMTP Email:", zap.Any("Mail Info", mailInfo))
	return nil
}

/**
 * @Description: 发送简单邮件
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
 * @Description: 批量发送邮件
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
 * @Description: 守护进程使用通道在窗口时间内批量发送邮件
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
