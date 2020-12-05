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

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
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

	db, err = NewORMDb(config)
	return err
}
