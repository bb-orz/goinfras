package mysqlStore

import (
	"database/sql"
	"fmt"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewMysqlClient(config *mysqlConfig) (db *sql.DB, err error) {
	db, err = manager.New(config.DbName, config.DbUser, config.DbPasswd, config.DbHost).Set(
		manager.SetCharset(config.ChartSet),                                   // 设置编码类型：utf8
		manager.SetAllowCleartextPasswords(config.AllowCleartextPasswords),    // 开发环境中设置允许明文密码通信
		manager.SetInterpolateParams(config.InterpolateParams),                // 设置允许占位符参数
		manager.SetTimeout(time.Duration(config.Timeout)*time.Second),         // 连接超时时间
		manager.SetReadTimeout(time.Duration(config.ReadTimeout)*time.Second), // 读超时时间
		manager.SetParseTime(config.ParseTime),                                // 将数据库的datetime时间格式转换为go time包数据类型
	).Port(int(config.DbPort)).Open(config.PING)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second) // 设置最大的连接时间，1分钟
	db.SetMaxIdleConns(int(config.MaxIdleConns))                               // 设置最大的闲置连接数
	db.SetMaxOpenConns(int(config.MaxOpenConns))                               // 设置最大的连接数
	fmt.Println("Mysql pool init ready!")

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Mysql Connect Successful!")
	}

	return db, nil
}
