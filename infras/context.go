package infras
import (
	"github.com/tietang/props/kvs"
)

const (
	KeyConfig = "_conf"
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
