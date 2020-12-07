package XGorm

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// 创建一个默认配置的DB
func CreateDefaultDB(config *Config) error {
	var err error
	if config == nil {
		config = DefaultConfig()
	}
	db, err = NewORMDb(config)
	return err
}

func XDB() *gorm.DB {
	return db
}

// 资源组件闭包执行
func XFDB(f func(c *gorm.DB) error) error {
	return f(db)
}
