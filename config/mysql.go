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
	db, err := sql.Open("mysql", "admin:server12321@tcp(beter-test-db.cj1xvqwlwiwo.ap-northeast-2.rds.amazonaws.com:3306)/outframe")
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
