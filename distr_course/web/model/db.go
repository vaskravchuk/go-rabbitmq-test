package model

import (
	"database/sql"
	_ "github.com/lib/pq"
)


var db *sql.DB


func init() {
	var err error
	db, err = sql.Open(
		"postgres",
		"postgres://postgres:123@10.0.1.13/go-rabbitmq-test?sslmode=disable")

	if err != nil {
		panic(err.Error())
	}
}