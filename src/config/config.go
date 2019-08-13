package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// GetDB digunakan untuk koneksi database
func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbHost := "remotemysql.com"
	dbUser := "2XstgudjDO"
	dbPass := "hxB1bNpYV9"
	dbName := "2XstgudjDO"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName)
	return
}
