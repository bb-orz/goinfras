package XSQLBuilder

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	_ "github.com/go-sql-driver/mysql"
)

type starter struct {
	goinfras.BaseStarter
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
func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Mysql", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	s.cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v ", define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	db, err = NewDB(s.cfg)
	if sctx.PassError(s.Name(), goinfras.StepSetup, err) {
		sctx.Logger().Info(s.Name(), goinfras.StepSetup, "SQLBuilder DB Setuped! ")
	}
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(db)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().OK(s.Name(), goinfras.StepCheck, "SQLBuilder DB Setup Successful! ")
		return true
	}
	return false
}

func (s *starter) Stop() {
	db.Close()
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.ResourcesGroup }
