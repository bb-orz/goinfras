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
	if err != nil {
		panic(err.Error())
	}
	s.cfg = &define
}

// 检查该组件的前置依赖
func (s *MysqlStarter) Setup(sctx *infras.StarterContext) {}

// 启动该资源组件
func (s *MysqlStarter) Start(sctx *infras.StarterContext) {
	var err error
	mysqlClient,err = NewMysqlClient(s.cfg)
	if err != nil {
		panic(err.Error())
	}
}

// 停止服务
func (s *MysqlStarter) Stop(sctx *infras.StarterContext)  {

}