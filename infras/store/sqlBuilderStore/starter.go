package sqlBuilderStore

import (
	"GoWebScaffold/infras"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Starter struct {
	infras.BaseStarter
	cfg Config
}

// 读取配置
func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Mysql", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	var d *sql.DB
	d, err = NewDB(&s.cfg)
	infras.FailHandler(err)
	SetComponent(d)
	sctx.Logger().Info("MysqlClient Setup Successful ...")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	SqlBuilderComponent().Close()
}
