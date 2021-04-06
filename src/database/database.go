package database

import (
	"challenge-golang-stone/src/config"
	"database/sql"
)

// Connect will return a db connection or an error
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
