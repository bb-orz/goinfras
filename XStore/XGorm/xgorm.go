package XGorm

import (
	"gorm.io/gorm"
)

func XDB() *gorm.DB {
	return db
}

// 资源组件闭包执行
func XFDB(f func(c *gorm.DB) error) error {
	return f(db)
}
