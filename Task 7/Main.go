package main

import (
	_ "pq-master"

	"./db"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbName   = "Task7"
)

func main() {
	dbase, err := db.Open()
}
