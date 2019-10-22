package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" //Driver for postgresql
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "alex"
	dbname   = "task8"
)

func Open() (*sql.DB, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", info)
	if err != nil {
		panic(err)
	}
	return db, err
}

func startRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/getPhoneNums", getPhoneNums).Methods("GET")
	fmt.Println("Started serving port :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	startRouter()
}

func getPhoneNums(w http.ResponseWriter, r *http.Request) {
	phoneNums := getPhoneNumsFromDb()
	json.NewEncoder(w).Encode(phoneNums)
}

func getPhoneNumsFromDb() []string {
	dbase, err := Open()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := dbase.Query("select \"PhoneNum\" from public.\"phone_numbers\" ")
	phoneNums := make([]string, 0, 8)
	for rows.Next() {
		phoneNum := ""
		err := rows.Scan(&phoneNum)
		if err != nil {
			fmt.Println(err)
			continue
		}
		phoneNums = append(phoneNums, phoneNum)
	}
	return phoneNums
}
