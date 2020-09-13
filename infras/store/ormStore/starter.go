package ormStore

import (
	"GoWebScaffold/infras"
	"github.com/jinzhu/gorm"
)

var gormDb *gorm.DB

func GormDb() *gorm.DB {
	infras.Check(gormDb)
	return gormDb
}

type ORMStarter struct {
	infras.BaseStarter
	cfg *OrmConfig
}

// 读取配置
func (s *ORMStarter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := OrmConfig{}
	err := viper.UnmarshalKey("OrmConfig", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

// 连接数据库
func (s *ORMStarter) Setup(sctx *infras.StarterContext) {
	var err error
	gormDb, err = NewORMDb(s.cfg)
	infras.FailHandler(err)
	sctx.Logger().Info("GormClient Setup Successful ...")
}

func (s *ORMStarter) Stop(sctx *infras.StarterContext) {
	GormDb().Close()
}

// 测试时启动db连接
func RunForTesting(config *OrmConfig) error {
	var err error
	if config == nil {
		config = &OrmConfig{
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

	gormDb, err = NewORMDb(config)
	return err
}
