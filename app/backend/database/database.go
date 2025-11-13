package database

import (
	"database/sql"
	"strings"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connstr := "host=localhost port=5432 user=postgres password=tuogjl100 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		connstr := "host=localhost port=5433 user=postgres password=123 dbname=postgres sslmode=disable"
		db, err = sql.Open("postgres", connstr)
		if err != nil {
			return nil, err
		}
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

func CreateDatabase(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE myapp")
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
	}
	return nil
}
