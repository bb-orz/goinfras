package mysqlStore

import (
	"GoWebScaffold/infras"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tietang/props/kvs"
)

var mysqlClient *sql.DB

func MysqlClient() *sql.DB {
	infras.Check(mysqlClient)
	return mysqlClient
}

type MysqlStarter struct {
	infras.BaseStarter
	cfg *mysqlConfig
}

// 读取配置
func (s *MysqlStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := mysqlConfig{}
	err := kvs.Unmarshal(configs, &define, "Mysql")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *MysqlStarter) Setup(sctx *infras.StarterContext) {
	var err error
	mysqlClient, err = NewMysqlClient(s.cfg)
	infras.FailHandler(err)
	sctx.Logger().Info("MysqlClient Setup Successful ...")
}

func (s *MysqlStarter) Stop(sctx *infras.StarterContext) {
	MysqlClient().Close()
}
