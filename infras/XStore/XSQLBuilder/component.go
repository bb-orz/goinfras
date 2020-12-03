package XSQLBuilder

import (
	"GoWebScaffold/infras"
	"database/sql"
)

var db *sql.DB

func SqlBuilderComponent() *sql.DB {
	infras.Check(db)
	return db
}

func SetComponent(d *sql.DB) {
	db = d
}
