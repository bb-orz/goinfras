package XSQLBuilder

import (
	"database/sql"
)

var db *sql.DB

func XDB() *sql.DB {
	return db
}

// 资源组件闭包执行
func XFDB(f func(c *sql.DB) error) error {
	return f(db)
}

// sqlbuilder通用操作实例
func XCommon() *CommonDao {
	dao := new(CommonDao)
	dao.db = XDB()
	return dao
}

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			"127.0.0.1",
			3306,
			"",
			"",
			"",
			60,
			100,
			200,
			"uft8",
			true,
			true,
			5,
			30,
			true,
			true,
		}
	}
	db, err = NewDB(config)
	return err
}
