package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

// CloseDB closes the database connection
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatal(err)
	}
}

// Exec executes a query without returning any rows.
func Exec(query string, args ...interface{}) (sql.Result, error) {
	return DB.Exec(query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
func QueryRow(query string, args ...interface{}) *sql.Row {
	return DB.QueryRow(query, args...)
}

// Query executes a query that returns multiple rows.
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return DB.Query(query, args...)
}
