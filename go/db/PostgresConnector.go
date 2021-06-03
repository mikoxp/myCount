package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func CreateCon() *sql.DB {

	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, database_name)

	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		log.Println(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		log.Println("Postgres db is not connected")
		log.Println(err.Error())
	}
	return db
}
