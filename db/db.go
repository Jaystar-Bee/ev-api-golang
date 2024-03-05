package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() {

	DB, err := sql.Open("sqlite3", "event_api.db")

	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(15)
	DB.SetMaxIdleConns(5)
}
