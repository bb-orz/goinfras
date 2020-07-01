package global

type GlobalConfig struct {
	Env        string `val:"debug"`
	AppName    string `val:"My App"`
	ServerName string `val:"example.com"`
}
