package main

import (
	"database/sql"
	"log"
)

// DB - Database connection object
var DB *sql.DB

// InitDatabase - initialize database connection
func InitDatabase() {
	db, err := sql.Open("sqlite3", "template.db")
	if err != nil {
		log.Fatalln(err)
	}

	// Verify connection
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	DB = db
}

// InitMigrations - create tables
func InitMigrations() {
	createTableUser := `CREATE TABLE IF NOT EXISTS users (
												id INTEGER PRIMARY KEY,
												username TEXT NOT NULL,
												password TEXT NOT NULL,
												name TEXT NOT NULL,
												email TEXT NOT NULL,
												phone TEXT NOT NULL
											);`
	_, err := DB.Exec(createTableUser)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
