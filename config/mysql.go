package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrDBConnection = errors.New("DB has something problem")
)

func ConnectionMysql() (*sqlx.DB, error) {
	str := makeMySQLSource()

	db, err := sqlx.Open("mysql", str)
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

func makeMySQLSource() string {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	if user == "" || pass == "" {
		log.Fatalln("You must include MYSQL_USER and MYSQL_PASSWORD environment variables")
	}

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DATABASE")

	if host == "" || port == "" || dbname == "" {
		log.Fatalln("You must include MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE environment variables")
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)
}
