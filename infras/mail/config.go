package mail

type MailConfig struct {
	NoAuth   bool   `val:"false"`              // 使用本地SMTP服务器发送电子邮件。
	NoSmtp   bool   `val:"false"`              // 使用API​​或后缀发送电子邮件。
	Server   string `val:"smtp.qq.com"`        // 使用外部SMTP服务器
	Port     int    `val:"587"`                // 外部SMTP服务端口
	User     string `val:"your email address"` // 你的三方邮箱地址
	Password string `val:""`                   // 你的邮箱密码
}
