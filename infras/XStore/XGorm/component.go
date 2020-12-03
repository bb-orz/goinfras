package XGorm

import (
	"GoWebScaffold/infras"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func ORMComponent() *gorm.DB {
	infras.Check(db)
	return db
}

func SetComponent(d *gorm.DB) {
	db = d
}
