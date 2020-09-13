package mail

type MailConfig struct {
	NoAuth   bool   // 使用本地SMTP服务器发送电子邮件。
	NoSmtp   bool   // 使用API​​或后缀发送电子邮件。
	Server   string // 使用外部SMTP服务器
	Port     int    // 外部SMTP服务端口
	User     string // 你的三方邮箱地址
	Password string // 你的邮箱密码
}
