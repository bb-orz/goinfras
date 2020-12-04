package XSQLBuilder

import (
	"GoWebScaffold/infras"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type starter struct {
	infras.BaseStarter
	cfg Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = Config{}
	return starter
}

func (s *starter) Name() string {
	return "XSQLBuilder"
}

// 读取配置
func (s *starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Mysql", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	var err error
	db, err = NewDB(&s.cfg)
	infras.FailHandler(err)
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(db)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: SQL Builder DB Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: SQL Builder DB Setup Successful!", s.Name()))
	return true
}

func (s *starter) Stop(sctx *infras.StarterContext) {
	db.Close()
}
