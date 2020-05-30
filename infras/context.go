package infras

import (
	"github.com/tietang/props/kvs"
	"go.uber.org/zap"
)

const (
	KeyConfig = "_conf"
	KeyLogger = "_logger"
)

//资源启动器上下文，
// 用来在服务资源初始化、安装、启动和停止的生命周期中变量和对象的传递
type StarterContext map[string]interface{}

func (s StarterContext) Configs() kvs.ConfigSource {
	p := s[KeyConfig]
	if p == nil {
		panic("配置还没有被初始化")
	}
	return p.(kvs.ConfigSource)
}
func (s StarterContext) SetConfigs(conf kvs.ConfigSource) {
	s[KeyConfig] = conf
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
