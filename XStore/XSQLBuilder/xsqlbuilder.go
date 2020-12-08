package XSQLBuilder

import (
	"database/sql"
	"goinfras"
)

var db *sql.DB

// 创建一个默认配置的Manager
func CreateDefaultDB(config *Config) error {
	var err error
	if config == nil {
		config = DefaultConfig()
	}
	db, err = NewDB(config)
	return err
}

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
