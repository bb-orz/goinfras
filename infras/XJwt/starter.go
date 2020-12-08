package XJwt

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/XStore/XRedis"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type starter struct {
	infras.BaseStarter
	cfg *Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	return starter
}

func (s *starter) Name() string {
	return "XJWT"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	var err error
	var define *Config

	// 先从viper读取配置信息
	viperConfig := sctx.Configs()
	if viperConfig != nil {
		err = viper.UnmarshalKey("Jwt", &define)
		infras.ErrorHandler(err)
	}

	// Viper读取不到配置时，default设置
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print Jwt Config:", zap.Any("JwtConfig", *define))
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	// 如果redis 连接池组件已安装，则缓存token到redis服务器
	if XRedis.CheckPool() {
		tku = NewTokenUtilsX(s.cfg)
	} else {
		tku = NewTokenUtils(s.cfg)
	}
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(tku)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: JWT TokenUtils Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: JWT TokenUtils Setup Successful!", s.Name()))
	return true
}

func (s *starter) Start(sctx *infras.StarterContext) {

}

func (s *starter) Stop() {}
