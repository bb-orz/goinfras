package starter

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/dao/RedisDao"
	"fmt"
)

type RedisStarter struct {
	infras.BaseStarter
}

func (s *RedisStarter) Init(sctx *StarterContext) {
	// 创建一个redis连接池
	pool := RedisDao.NewRedisPool(sctx.GetConfig())
	// 把连接池设置到上下文
	sctx.SetRedisPool(pool)
	fmt.Println("Redis Pool Client Init Successful!")
}
