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
	cfg *MysqlConfig
}

// 读取配置
func (s *MysqlStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := MysqlConfig{}
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

func RunForTesting(config *MysqlConfig) error {
	var err error
	if config == nil {
		config = &MysqlConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err = p.Unmarshal(config)
		if err != nil {
			return err
		}
	}
	mysqlClient, err = NewMysqlClient(config)
	return err
}
