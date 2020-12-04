package XGorm

import (
	"GoWebScaffold/infras"
	"fmt"
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
	return "XGorm"
}

// 读取配置
func (s *starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("OrmConfig", &define)
	infras.FailHandler(err)
	s.cfg = define
}

// 连接数据库
func (s *starter) Setup(sctx *infras.StarterContext) {
	var err error
	db, err = NewORMDb(&s.cfg)
	infras.FailHandler(err)
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(db)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: GORM DB Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: GORM DB Setup Successful!", s.Name()))
	return true
}

func (s *starter) Stop(sctx *infras.StarterContext) {
	db.Close()
}
