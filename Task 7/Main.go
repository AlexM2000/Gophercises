package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "pq-master"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbName   = "Task7"
)

func openDB() (*sql.DB, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbName=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err := sql.Open("postgres", info)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func main() {

}
