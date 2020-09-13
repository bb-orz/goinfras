package sqlbuilderStore

import (
	"GoWebScaffold/infras"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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
	viper := sctx.Configs()
	define := MysqlConfig{}
	err := viper.UnmarshalKey("Mysql", &define)
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
		config = &MysqlConfig{
			"127.0.0.1",
			3306,
			"",
			"",
			"",
			60,
			100,
			200,
			"uft8",
			true,
			true,
			5,
			30,
			true,
			true,
		}
	}
	sqlBuilderClient, err = NewMysqlClient(config)
	return err
}
