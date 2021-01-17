package goinfras

import "github.com/spf13/viper"

// 初始全局配置
const (
	Env      = "Env"      // 允许环境：dev、testing、product
	Host     = "Host"     // 主机地址
	Endpoint = "Endpoint" // 节点
	AppName  = "AppName"  // 应用名
	Version  = "Version"  // 应用版本
)

func NewGlobal(vpcfg *viper.Viper) Global {
	_g = make(map[string]interface{})
	_ = vpcfg.UnmarshalKey("Global", &_g)
	return _g
}

var _g Global

type Global map[string]interface{}

func (g Global) Set(k string, v interface{}) {
	g[k] = v
}

func (g Global) Get(k string) interface{} {
	if v, ok := g[k]; ok {
		return v
	}
	return nil
}

func (g Global) GetEnv() string {
	s := _g.Get(Env)
	if s == nil {
		return "undefined"
	}
	return s.(string)
}

func (g Global) GetHost() string {
	s := _g.Get(Host)
	if s == nil {
		return "undefined"
	}
	return s.(string)
}

func (g Global) GetEndpoint() string {
	s := _g.Get(Endpoint)
	if s == nil {
		return "undefined"
	}
	return s.(string)
}

func (g Global) GetAppName() string {
	s := _g.Get(AppName)
	if s == nil {
		return "undefined"
	}
	return s.(string)
}

func (g Global) GetVersion() string {
	s := _g.Get(Version)
	if s == nil {
		return "undefined"
	}
	return s.(string)
}
