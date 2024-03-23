package db

import (
	"backend-test-mkp/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.DB_HOST, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME, conf.DB_PORT))

	if err != nil {
		panic("Connection error ..." + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateConn() *sql.DB {
	return db
}
