package mail

import "gopkg.in/gomail.v2"

// 获取一个无鉴权的smtp服务拨号器
func NewNoAuthDialer(host string, port int) *gomail.Dialer {
	return &gomail.Dialer{Host: host, Port: port}
}

// 获取一个鉴权的smtp服务拨号器
func NewAuthDialer(host, user, password string, port int) *gomail.Dialer {
	return gomail.NewDialer(host, port, user, password)
}
