package XGorm

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

func NewORMDb(config *Config) (*gorm.DB, error) {
	var err error
	var sqlDBPool *sql.DB
	// 先创建一个*sql.DB 连接池
	sqlDBPool, err = newSqlDB(config)

	switch config.Dialect {
	case "mysql":
		db, err = gorm.Open(mysql.New(mysql.Config{
			Conn:                      sqlDBPool,
			DefaultStringSize:         config.DefaultStringSize,         // 字符串字段默认长度
			DisableDatetimePrecision:  config.DisableDatetimePrecision,  // 禁用datetime precision，这在MySQL5.6之前不受支持
			DontSupportRenameIndex:    config.DontSupportRenameIndex,    //  是否重命名索引，重命名索引在MySQL 5.7和MariaDB之前不支持
			DontSupportRenameColumn:   config.DontSupportRenameColumn,   // 是否可以重命名列，MySQL 8和MariaDB之前不支持重命名列
			SkipInitializeWithVersion: config.SkipInitializeWithVersion, // 是否基于当前MySQL版本自动配置
		}), &gorm.Config{})

	case "postgres":
		// source : "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword"
		db, err = gorm.Open(postgres.New(postgres.Config{
			Conn:                 sqlDBPool,
			PreferSimpleProtocol: true, // 是否禁用隐式 prepared 语法
		}), &gorm.Config{})

	default:
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	}

	return db, err
}
