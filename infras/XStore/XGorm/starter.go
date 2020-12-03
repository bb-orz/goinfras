package XGorm

import (
	"GoWebScaffold/infras"
	"github.com/jinzhu/gorm"
)

type Starter struct {
	infras.BaseStarter
	cfg Config
}

// 读取配置
func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("OrmConfig", &define)
	infras.FailHandler(err)
	s.cfg = define
}

// 连接数据库
func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	var d *gorm.DB
	d, err = NewORMDb(&s.cfg)
	infras.FailHandler(err)
	SetComponent(d)
	sctx.Logger().Info("GormClient Setup Successful ...")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	ORMComponent().Close()
}
