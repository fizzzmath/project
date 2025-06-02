package shared

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "u68867:6788851@/u68867")

	if err != nil {
		return nil, err
	}

	return db, nil
}