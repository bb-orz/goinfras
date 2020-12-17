package XEsOlivere

import (
	"context"
	"fmt"
	"github.com/bb-orz/goinfras"
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
		fmt.Printf("XEsOlivere Config: %v \n", *define)
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
	pingService := esClient.Ping(s.cfg.URL)
	result, _, err := pingService.Do(context.Background())
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Olivere ElasticSearch Client Ping Fail:%s", s.Name(), err.Error()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Olivere ElasticSearch Client Setup Successful! ES Server Info:[ServerName:%s,ClusterName:%s,TagLine:%s,Version:%v]", s.Name(), result.Name, result.ClusterName, result.TagLine, result.Version))
	return true
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.BasicGroup }
