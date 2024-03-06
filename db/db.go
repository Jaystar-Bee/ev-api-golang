package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() {
	Database, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}
	DB = Database

	DB.SetMaxOpenConns(15)
	DB.SetMaxIdleConns(5)

	createDatabaseTables()
}
