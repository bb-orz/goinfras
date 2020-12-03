package XGlobal

// 全局配置
var global Config

// 资源组件实例调用
func GConfig() *Config {
	return &global
}
