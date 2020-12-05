package XJwt

import "GoWebScaffold/infras/XStore/XRedis"

var tku ITokenUtils

// 资源组件实例调用
func XTokenUtils() ITokenUtils {
	return tku
}

// 资源组件闭包执行
func XFTokenUtils(f func(t ITokenUtils) error) error {
	return f(tku)
}

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error

	if config == nil {
		config = &Config{
			PrivateKey: "ginger_key",
			ExpSeconds: 60,
		}

	}
	tku = NewTokenUtils([]byte(config.PrivateKey), config.ExpSeconds)
	return err
}

/*实例化使用 redis cache 的资源用于测试*/
func TestingInstantiationForRedisCache(config *Config) error {
	var err error

	// 初始化依赖的redis缓存组件
	err = XRedis.TestingInstantiation()
	if err != nil {
		return err
	}

	if config == nil {
		config = &Config{
			PrivateKey: "ginger_key",
			ExpSeconds: 5,
		}

	}
	tku = NewTokenUtilsX([]byte(config.PrivateKey), config.ExpSeconds)
	return err
}
