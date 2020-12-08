package XSQLBuilder

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"goinfras"
)

type starter struct {
	BaseStarter
	cfg *Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	return starter
}

func (s *starter) Name() string {
	return "XSQLBuilder"
}

// 读取配置
func (s *starter) Init(sctx *StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Mysql", &define)
		ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print Mysql Config:", zap.Any("Mysql", *define))
}

func (s *starter) Setup(sctx *StarterContext) {
	var err error
	db, err = NewDB(s.cfg)
	FailHandler(err)
}

func (s *starter) Check(sctx *StarterContext) bool {
	err := Check(db)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: SQL Builder DB Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: SQL Builder DB Setup Successful!", s.Name()))
	return true
}

func (s *starter) Stop() {
	db.Close()
}