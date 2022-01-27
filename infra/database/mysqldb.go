package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDB(user, password, host string, port int, name string) (db *sql.DB, err error) {
	db, err = sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			user,
			password,
			host,
			port,
			name,
		),
	)

	if err != nil {
		return nil, fmt.Errorf("error while opening db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}

	return db, nil
}
