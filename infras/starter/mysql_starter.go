package starter

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/dao/mysqlDao"
	"fmt"
)

type MysqlStarter struct {
	infras.BaseStarter
}

func (s *MysqlStarter) Init(sctx *StarterContext) {
	// 创建一个mysql 连接客户端
	client := mysqlDao.NewMysqlClient(sctx.GetConfig())
	// 把客户端连接设置到上下文
	sctx.SetMysqlClient(client)
	fmt.Println("Mysql Client Init Successful!")
}
