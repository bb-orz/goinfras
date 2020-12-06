package infras

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	KeyConfig = "_vpcfg"
	KeyLogger = "_logger"
)

// 资源启动器上下文，
// 用来在服务资源初始化、安装、启动和停止的生命周期中变量和对象的传递
type StarterContext map[string]interface{}

func (s StarterContext) Configs() *viper.Viper {
	p := s[KeyConfig]
	if p == nil {
		panic("配置还没有被初始化")
	}
	return p.(*viper.Viper)
}
func (s StarterContext) SetConfigs(vpcfg *viper.Viper) {
	s[KeyConfig] = vpcfg
}

func (s StarterContext) Logger() *zap.Logger {
	p := s[KeyLogger]
	if p == nil {
		panic("日志记录器还没有被初始化")
	}
	return p.(*zap.Logger)
}
func (s StarterContext) SetLogger(logger *zap.Logger) {
	s[KeyLogger] = logger
}

// 创建一个默认最少配置启动器上下文
func CreateDefaultSystemContext() *StarterContext {
	sctx := &StarterContext{}
	sctx.SetConfigs(viper.New())
	sctx.SetLogger(zap.L())
	return sctx
}
