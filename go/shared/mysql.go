package shared

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "u68867:6788851@/u68867")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func UpdateUser(db *sql.DB, key string, value string, ID int) error {
	sql, err := db.Query(fmt.Sprintf(`
		UPDATE User
		SET %s = ?
		WHERE ID = ?
	`, key), value, ID)

	if err != nil {
		return err
	}

	defer sql.Close()

	return nil
}