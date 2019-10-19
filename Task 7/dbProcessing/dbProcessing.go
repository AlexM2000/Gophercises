package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mux-master"
	"net/http"
	_ "pq-master" //Driver for PostgreSQL
	"strconv"

	"./db"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "alex"
	dbname   = "Task7"
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

func StartRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/tasks/create", createTask).Methods("POST")
	r.HandleFunc("/tasks/complete/{id}", completeTask).Methods("PUT")
	fmt.Println("Started serving port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	StartRouter()
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	tasks := getTasksFromDb()
	json.NewEncoder(w).Encode(tasks)
}

func getTasksFromDb() []db.Task {
	dbase, err := Open()
	if err != nil {
		panic(err)
	}
	defer dbase.Close()
	Tasks := []db.Task{}

	rows, err := dbase.Query("SELECT * from public.\"Listtodo\" where \"DateComplete\" is null")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		t := db.Task{}
		err := rows.Scan(&t.Id, &t.Text, &t.CreateTime, &t.CompleteTime)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Tasks = append(Tasks, t)
	}
	return Tasks
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask db.Task
	json.NewDecoder(r.Body).Decode(&newTask)
	task := createTaskInDb(newTask)
	json.NewEncoder(w).Encode(task)
}

func createTaskInDb(task db.Task) db.Task {
	dbase, err := Open()
	if err != nil {
		log.Fatal(err)
	}
	_, err = dbase.Exec("insert into public.\"Listtodo\"(\"task\") values($1);", task.Text)
	if err != nil {
		log.Fatal(err)
	}
	row := dbase.QueryRow("select * from public.\"Listtodo\" where \"DateCreate\"=(select max(\"DateCreate\") from public.\"Listtodo\")")
	t := db.Task{}
	err = row.Scan(&t.Id, &t.Text, &t.CreateTime, &t.CompleteTime)
	fmt.Println("Scan success")
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func completeTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	var completeTask db.Task
	json.NewDecoder(r.Body).Decode(&completeTask)
	completeTask.Id = id
	completedTask := completeTaskInDb(completeTask)
	json.NewEncoder(w).Encode(completedTask)
}

func completeTaskInDb(task db.Task) db.Task {
	dbase, err := Open()
	if err != nil {
		log.Fatal(err)
	}
	defer dbase.Close()
	_, err = dbase.Exec("update public.\"Listtodo\" set \"DateComplete\" = CURRENT_TIMESTAMP where \"Id\" = $1", task.Id)
	if err != nil {
		log.Fatal(err)
	}
	row := dbase.QueryRow("select * from public.\"Listtodo\" where \"DateComplete\"=(select max(\"DateComplete\") from public.\"Listtodo\")")
	t := db.Task{}
	err = row.Scan(&t.Id, &t.Text, &t.CreateTime, &t.CompleteTime)
	fmt.Println("Scan success")
	if err != nil {
		fmt.Println(err)
	}
	return t
}
