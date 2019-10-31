package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"./table"
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
	r.HandleFunc("/get", getData).Methods("GET")
	r.HandleFunc("/getPhone/{phonenum}", findPhoneNum).Methods("GET")
	r.HandleFunc("/deletePhone/{id}", deletePhoneNum).Methods("DELETE")
	r.HandleFunc("/updatePhone/{id}", updatePhone).Methods("PUT")
	fmt.Println("Started serving port :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	startRouter()
}

func getData(w http.ResponseWriter, r *http.Request) {
	data := getDataFromDb()
	json.NewEncoder(w).Encode(data)
}

func findPhoneNum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tableRow := findPhoneNumInDb(params["phonenum"])
	json.NewEncoder(w).Encode(tableRow)
}

func updatePhone(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	t := table.Table{}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(r.Body).Decode(&t)
	t.ID = id
	tableRow := updatePhoneNumInDb(t)
	json.NewEncoder(w).Encode(tableRow)
}

func updatePhoneNumInDb(t table.Table) table.Table {
	dbase, err := Open()
	if err != nil {
		log.Fatal(err)
	}
	defer dbase.Close()
	_, err = dbase.Exec("update public.\"phone_numbers\" set \"PhoneNum\"=$1 where \"Id\"=$2", t.PhoneNum, t.ID)
	if err != nil {
		fmt.Println("err")
		log.Fatal(err)
	}
	return t
}

func deletePhoneNum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	t := table.Table{}
	json.NewDecoder(r.Body).Decode(&t)
	t.ID = id
	deletePhoneFromDb(t)
	json.NewEncoder(w).Encode("DELETE OK")
}

func deletePhoneFromDb(t table.Table) {
	dbase, err := Open()
	if err != nil {
		log.Fatal(err)
	}
	defer dbase.Close()
	_, err = dbase.Exec("delete from public.\"phone_numbers\" where \"Id\"=$1;", t.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func findPhoneNumInDb(number string) table.Table {
	dbase, err := Open()
	if err != nil {
		log.Fatal(err)
	}
	defer dbase.Close()
	var t table.Table
	row := dbase.QueryRow("select * from public.\"phone_numbers\" where \"PhoneNum\"=$1", number)
	err = row.Scan(&t.PhoneNum, &t.FirstName, &t.SecondName, &t.ThirdName, &t.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			t.PhoneNum = ""
		} else {
			fmt.Println(err)
		}
	}
	return t
}

func getDataFromDb() []table.Table {
	dbase, err := Open()
	if err != nil {
		log.Fatal(err)
	}
	defer dbase.Close()
	rows, err := dbase.Query("select distinct * from public.\"phone_numbers\" ")
	tableRows := make([]table.Table, 0, 8)
	for rows.Next() {
		tableRow := table.Table{}
		err := rows.Scan(&tableRow.PhoneNum, &tableRow.FirstName,
			&tableRow.SecondName, &tableRow.ThirdName, &tableRow.ID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tableRows = append(tableRows, tableRow)
	}
	return tableRows
}
