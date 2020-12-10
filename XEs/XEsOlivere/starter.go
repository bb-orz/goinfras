package XEsOlivere

import (
	"fmt"
	"go.uber.org/zap"
	"goinfras"
)

// 实例化资源存储变量

/* 资源启动器 */
type starter struct {
	goinfras.BaseStarter
	cfg *Config
}

// 应用注册启动器时创建
func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	return starter
}

func (s *starter) Name() string {
	return "XEs"
}

// 应用初始化时加载配置数据
func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("XEsOlivere", &define)
		goinfras.ErrorHandler(err)
	}
	if define != nil {
		sctx.Logger().Info("Print XEsOlivere Config:", zap.Any("XEsOlivereConfig", *define))
	}
	s.cfg = define
}

// 应用安装阶段创建Cron管理器，并注册为应用组件
func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	esClient, err = NewESClient(s.cfg)
	goinfras.ErrorHandler(err)
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(esClient)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Olivere ElasticSearch Client Setup Fail!", s.Name()))
		return false
	}

	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Olivere ElasticSearch Client Setup Successful!", s.Name()))
	return true

}
