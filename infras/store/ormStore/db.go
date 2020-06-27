package ormStore

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewORMDb(config *ormConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	switch config.Dialect {
	case "mysql":
		// source : "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
		source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
			config.DbUser,
			config.DbPasswd,
			config.DbHost,
			config.DbPort,
			config.DbName,
			config.ChartSet,
			config.ParseTime,
			config.Local,
		)
		db, err = gorm.Open("mysql", source)

	case "postgres":
		// source : "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword"
		source := fmt.Sprintf("host=%s:%d user=%s dbname=%s password=%s sslmode=%s",
			config.DbHost,
			config.DbPort,
			config.DbUser,
			config.DbName,
			config.DbPasswd,
			config.SSLMode,
		)
		db, err = gorm.Open("postgres", source)
	default:
		return nil, errors.New("sql driver is invalid! ")
	}

	if db != nil {
		db.SingularTable(config.SingularTable)
		err := db.DB().Ping()
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
