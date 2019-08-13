package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// GetDB digunakan untuk koneksi database
func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbHost := "remotemysql.com"
	dbUser := "xxxxx" //input your username from remotemysql
	dbPass := "xxxxx" //input your password from remotemysql
	dbName := "xxxxx" //input your databasename from remotemysql
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName)
	return
}
