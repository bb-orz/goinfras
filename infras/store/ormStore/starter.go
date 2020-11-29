package ormStore

import (
	"GoWebScaffold/infras"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Db() *gorm.DB {
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
	err := viper.UnmarshalKey("OrmConfig", &define)
	infras.FailHandler(err)
	s.cfg = define
}

// 连接数据库
func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	db, err = NewORMDb(&s.cfg)
	infras.FailHandler(err)
	sctx.Logger().Info("GormClient Setup Successful ...")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	Db().Close()
}

// 测试时启动db连接
func RunForTesting(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			"mysql",
			"127.0.0.1",
			3306,
			"dev",
			"123456",
			"dev_db",
			"utf8",
			true,
			"Local",
			"disable",
			false,
		}
	}

	db, err = NewORMDb(config)
	return err
}
