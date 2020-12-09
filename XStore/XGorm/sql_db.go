package XGorm

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	_ "github.com/go-sql-driver/mysql"
)

// 创建一个sql.DB的mysql数据库连接
func newSqlDB(config *Config) (*sql.DB, error) {
	var err error
	var db *sql.DB
	switch config.Dialect {
	case "mysql":
		mysqlDSN := fmt.Sprintf("gorm:gorm@tcp(%s:%d)/gorm?charset=%s&parseTime=%t&loc=%s", config.DbHost, config.DbPort, config.ChartSet, config.ParseTime, config.Local)
		db, err = sql.Open("mysql", mysqlDSN)
	case "postgres":
		postgresDSN := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", config.DbUser, config.DbPasswd, config.DbName, config.DbPort, config.SslMode, config.TimeZone)
		db, err = sql.Open("postgres", postgresDSN)
	}

	if db != nil {
		db.SetConnMaxLifetime(config.ConnMaxLifetime) // 设置了连接可复用的最大时间
		db.SetMaxIdleConns(config.MaxIdleConns)       // 设置连接池中空闲连接的最大数量
		db.SetMaxOpenConns(config.MaxOpenConns)       // 设置打开数据库连接的最大数量

		err = db.Ping()
		if err != nil {
			return nil, err
		}
	}

	return db, err
}
