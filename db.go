package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./shrink.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS urls(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			short_code TEXT NOT NULL,
			original_url TEXT NOT NULL
		)
	`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to the SQLite database successfully.")

}
