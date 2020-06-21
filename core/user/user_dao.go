package user

import (
	"GoWebScaffold/infras/store/mysqlStore"
)

// 数据访问层，实现具体数据持久化操作
type UserDAO struct {
	dao *mysqlStore.BaseDao
}
