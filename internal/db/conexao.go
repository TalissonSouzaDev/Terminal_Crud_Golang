package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conn(config string) *sql.DB {
	conn, err := sql.Open("mysql", config)
	if err != nil {
		panic(err)
	}
	return conn
}
