package global

type Config struct {
	Env        string // 允许环境：dev、testing、product
	AppName    string // 应用名称
	ServerName string // 服务名称
}
