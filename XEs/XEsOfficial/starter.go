package XEsOfficial

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/elastic/go-elasticsearch/v8/estransport"
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
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("EsOfficial", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	s.cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v ", define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	if s.optCfg == nil {
		s.optCfg = &OptionalConfig{}
	}
	esClient, err = NewESClient(s.cfg, s.optCfg.HttpHeader, s.optCfg.HttpTransport, s.optCfg.Logger, s.optCfg.Selector, s.optCfg.RetryBackoffFunc, s.optCfg.ConnectionPoolFunc)
	if sctx.PassError(s.Name(), goinfras.StepSetup, err) {
		sctx.Logger().Info(s.Name(), goinfras.StepSetup, "Es Official Client Setuped! ")
	}
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(esClient)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().OK(s.Name(), goinfras.StepCheck, "Es Official Client Setup Successful! ")
	}

	_, err = esClient.Ping()
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().OK(s.Name(), goinfras.StepCheck, "Es Official Client Ping Successful! ")
		return true
	}
	return false
}

// 设置启动组级别:
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.BasicGroup }
