package ormStore

import (
	"GoWebScaffold/infras"
	"github.com/jinzhu/gorm"
	"github.com/tietang/props/kvs"
)

var gormDb *gorm.DB

func GormDb() *gorm.DB {
	infras.Check(gormDb)
	return gormDb
}

type MysqlStarter struct {
	infras.BaseStarter
	cfg *OrmConfig
}

// 读取配置
func (s *MysqlStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := OrmConfig{}
	err := kvs.Unmarshal(configs, &define, "OrmConfig")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *MysqlStarter) Setup(sctx *infras.StarterContext) {
	var err error
	gormDb, err = NewORMDb(s.cfg)
	infras.FailHandler(err)
	sctx.Logger().Info("GormClient Setup Successful ...")
}

func (s *MysqlStarter) Stop(sctx *infras.StarterContext) {
	GormDb().Close()
}
