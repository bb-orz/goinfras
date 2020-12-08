package XJwt

import (
	"GoWebScaffold/infras/XStore/XRedis"
	"go.uber.org/zap"
)

var tku ITokenUtils

// 创建一个默认配置的TokenUtils
func CreateDefaultTku(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	tku = NewTokenUtils(config)
}

// 创建一个默认配置的带redis缓存的TokenUtils
func CreateDefaultTkuX(config *Config) error {
	if config == nil {
		config = DefaultConfig()
	}

	// 检查redis连接池组件或创建默认池
	if !XRedis.CheckPool() {
		logger, err := zap.NewDevelopment()
		if err != nil {
			return err
		}
		err = XRedis.CreateDefaultPool(nil, logger)
		if err != nil {
			return err
		}
	}

	tku = NewTokenUtilsX(config)

	return nil
}

// 资源组件实例调用
func XTokenUtils() ITokenUtils {
	return tku
}

// 资源组件闭包执行
func XFTokenUtils(f func(t ITokenUtils) error) error {
	return f(tku)
}
