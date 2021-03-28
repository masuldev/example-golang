package config

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrDBConnection = errors.New("DB has something problem")
)

func ConnectionMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/test-db")
	if err != nil {
		return nil, ErrDBConnection
	}

	err = db.Ping()
	if err != nil {
		return nil, ErrDBConnection
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)

	return db, nil
}
