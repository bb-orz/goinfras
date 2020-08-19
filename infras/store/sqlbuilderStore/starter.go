package sqlbuilderStore

import (
	"GoWebScaffold/infras"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tietang/props/kvs"
)

var sqlBuilderClient *sql.DB

func SqlBuilderClient() *sql.DB {
	infras.Check(sqlBuilderClient)
	return sqlBuilderClient
}

type SqlBuilderStarter struct {
	infras.BaseStarter
	cfg *MysqlConfig
}

// 读取配置
func (s *SqlBuilderStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := MysqlConfig{}
	err := kvs.Unmarshal(configs, &define, "Mysql")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *SqlBuilderStarter) Setup(sctx *infras.StarterContext) {
	var err error
	sqlBuilderClient, err = NewMysqlClient(s.cfg)
	infras.FailHandler(err)
	sctx.Logger().Info("MysqlClient Setup Successful ...")
}

func (s *SqlBuilderStarter) Stop(sctx *infras.StarterContext) {
	SqlBuilderClient().Close()
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
	sqlBuilderClient, err = NewMysqlClient(config)
	return err
}
