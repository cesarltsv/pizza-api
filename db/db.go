package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	newDB, err := sql.Open("sqlite3", "./api.db")
	DB = newDB
	if err != nil {
		panic("Could not connect to database.")
	}
	defer DB.Close()

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createPizzasTable := `
	CREATE TABLE IF NOT EXISTS pizzas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		ingredients TEXT NOT NULL,
		createOn DATETIME NOTE NULL,
		user_id INTEGER

	)
	`
	_, err := DB.Exec(createPizzasTable)

	if err != nil {
		panic("Could not create pizza table")
	}
}
