package XEsOfficial

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/elastic/go-elasticsearch/v8/estransport"
	"go.uber.org/zap"
	"net/http"
)

// 实例化资源存储变量

/* 资源启动器 */
type starter struct {
	goinfras.BaseStarter
	cfg    *Config
	optCfg *OptionalConfig
}

// 可选配置设置
type OptionalConfig struct {
	HttpHeader         http.Header          // 设置API HTTP Header
	HttpTransport      http.RoundTripper    // 设置API HTTP transport object
	Logger             estransport.Logger   // 设置logger object
	Selector           estransport.Selector // 设置selector object
	RetryBackoffFunc   RetryBackoffFunc     // 设置可选的退避持续时间处理函数
	ConnectionPoolFunc ConnectionPoolFunc   // 设置连接池处理函数
}

// 应用注册启动器时创建
func NewStarter(optCfg *OptionalConfig) *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	starter.optCfg = optCfg
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
		err = viper.UnmarshalKey("XEsOfficial", &define)
		goinfras.ErrorHandler(err)
	}
	if define != nil {
		sctx.Logger().Info("Print XEsOfficial Config:", zap.Any("XEsOfficialConfig", *define))
	}
	s.cfg = define
}

// 应用安装阶段创建Cron管理器，并注册为应用组件
func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	if s.optCfg == nil {
		s.optCfg = &OptionalConfig{}
	}
	esClient, err = NewESClient(s.cfg, s.optCfg.HttpHeader, s.optCfg.HttpTransport, s.optCfg.Logger, s.optCfg.Selector, s.optCfg.RetryBackoffFunc, s.optCfg.ConnectionPoolFunc)
	goinfras.ErrorHandler(err)
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(esClient)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Official ElasticSearch Client Setup Fail!", s.Name()))
		return false
	}

	_, err = esClient.Ping()
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Official ElasticSearch Client Ping Error:%s", s.Name(), err.Error()))
		return false
	} else {
		sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Official ElasticSearch Client Setup Successful!", s.Name()))
		return true
	}

}

// 设置启动组级别:
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.BasicGroup }
