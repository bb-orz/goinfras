package sqlBuilderStore

import (
	"GoWebScaffold/infras"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DB() *sql.DB {
	infras.Check(db)
	return db
}

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
	db, err = NewDB(&s.cfg)
	infras.FailHandler(err)
	sctx.Logger().Info("MysqlClient Setup Successful ...")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	DB().Close()
}

func RunForTesting(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
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
	db, err = NewDB(config)
	return err
}
