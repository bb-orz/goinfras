package XGlobal

// 全局配置
const (
	Env      = "Env"      // 允许环境：dev、testing、product
	Host     = "Host"     // 主机地址
	Endpoint = "Endpoint" // 节点
	AppName  = "AppName"  // 应用名
	Version  = "Version"  // 应用版本
)

// Usage: env := XGlobal.G().Get("your global config key")
func G() Global {
	return _g
}

func GetEnv() string {
	s := _g.Get(Env)
	if s == nil {
		return "undefined"
	}
	return s.(string)
}

func GetHost() string {
	s := _g.Get(Host)
	if s == nil {
		return "undefined"
	}
	return s.(string)
}

func GetEndpoint() string {
	s := _g.Get(Endpoint)
	if s == nil {
		return "undefined"
	}
	return s.(string)
}

func GetAppName() string {
	s := _g.Get(AppName)
	if s == nil {
		return "undefined"
	}
	return s.(string)
}

func GetVersion() string {
	s := _g.Get(Version)
	if s == nil {
		return "undefined"
	}
	return s.(string)
}
