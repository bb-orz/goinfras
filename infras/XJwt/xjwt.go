package XJwt

var tku ITokenUtils

// 创建一个默认配置的TokenUtils
func CreateDefaultTku(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	tku = NewTokenUtils(config.PrivateKey, config.ExpSeconds)
}

// 创建一个默认配置的带redis缓存的TokenUtils
func CreateDefaultTkuX(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	tku = NewTokenUtilsX(config.PrivateKey, config.ExpSeconds)

	// TODO 创建一个redis 连接实例
}

// 资源组件实例调用
func XTokenUtils() ITokenUtils {
	return tku
}

// 资源组件闭包执行
func XFTokenUtils(f func(t ITokenUtils) error) error {
	return f(tku)
}
