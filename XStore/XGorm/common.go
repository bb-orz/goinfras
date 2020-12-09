package XGorm

import (
	"gorm.io/gorm"
)

type CommonGORM struct {
	db *gorm.DB
}

func XCommon() *CommonGORM {
	common := new(CommonGORM)
	common.db = XDB()
	return common
}

// TODO 定义一些通用操作，参考run_test.go
