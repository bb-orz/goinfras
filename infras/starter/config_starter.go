package starter

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/config"
)

type ConfigStarter struct {
	infras.BaseStarter
}

// 启动器初始化
func (s *ConfigStarter) Init(sctx *StarterContext) {
	// 启动时先加载配置文件并解析,设置解析后的配置信息到应用启动器上下文
	sctx.SetConfig(config.Parse())
}
