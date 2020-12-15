package XMail

import "gopkg.in/gomail.v2"

var mailDialer *gomail.Dialer

// 创建一个默认配置的Manager
func CreateDefaultManager(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	mailDialer = NewNoAuthDialer(config.Server, config.Port)
}

// 创建一个默认配置的Manager
func CreateDefaultAuthManager(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	mailDialer = NewAuthDialer(config.Server, config.User, config.Password, config.Port)
}

// 获取一个无鉴权的smtp服务拨号器
func NewNoAuthDialer(host string, port int) *gomail.Dialer {
	return &gomail.Dialer{Host: host, Port: port}
}

// 获取一个鉴权的smtp服务拨号器
func NewAuthDialer(host, user, password string, port int) *gomail.Dialer {
	return gomail.NewDialer(host, port, user, password)
}
