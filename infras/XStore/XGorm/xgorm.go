package XGorm

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func XDB() *gorm.DB {
	return db
}

// 资源组件闭包执行
func XFDB(f func(c *gorm.DB) error) error {
	return f(db)
}
