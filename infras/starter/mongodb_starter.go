package starter

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/dao/mongoDao"
	"fmt"
)

type MongoDBStarter struct {
	infras.BaseStarter
}

func (s *MongoDBStarter) Init(sctx *StarterContext) {
	// 创建一个mongo session 连接池
	client := mongoDao.NewMongoClient(sctx.GetConfig())
	// 把连接池设置到上下文
	sctx.SetMongoClient(client)
	fmt.Println("MongoDB Client Init Successful!")
}
