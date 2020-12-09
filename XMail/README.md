# Gomail Starter

> 基于 gopkg.in/gomail.v2 包

### Documentation

> Documentation https://godoc.org/gopkg.in/gomail.v2



### Starter Usage
```
goinfras.RegisterStarter(XMail.NewStarter())

```

### XMail Config Setting

```
NoAuth   bool   // 使用本地SMTP服务器发送电子邮件。
NoSmtp   bool   // 使用API​​或后缀发送电子邮件。
Server   string // 使用外部SMTP服务器
Port     int    // 外部SMTP服务端口
User     string // 你的三方邮箱地址
Password string // 你的邮箱密码
```

### XMail  Usage

1.使用外部API发送邮件

```
err := XCommonMail().SendMailNoSMTP("from", "subject", "body", "bodyType", []string{"to1","to2"},func(from string, to []string, msg io.WriterTo) error {
    fmt.Println("From:", from)
    fmt.Println("To:", to)
    fmt.Println("Msg:", msg)

    // TODO 通过外部API发送邮件

    return nil
})


```

2.使用SMTP服务器，发送简单邮件，测试前请先设置默认配置信息

```

// 发送
err := XCommonMail().SendSimpleMail(
    "from",
    "ccAddress",
    "ccName",
    "subject",
    "body",
    "bodyType",
    "attach",
    []string{"to1","to2"},
)
// ...
```


3.使用SMTP服务器，群发邮件，测试前请先设置默认配置信息
```
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

err := XCommonMail().SendNewsLetter(receivers, "frome", "subject", "body", "bodyType")


```


4.使用SMTP服务器，使用通道在窗口时间内批量发送邮件

```
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


```