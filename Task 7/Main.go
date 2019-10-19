package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"./db"
)

func main() {
	//fmt.Println(getTasks())
	//fmt.Println(createTask("To write program"))
	fmt.Println(completeTask(30))
}

func getTasks() []db.Task {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", "http://localhost:8000/tasks", nil,
	)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var tasks []db.Task
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bodyBytes, &tasks)
		if err != nil {
			panic(err)
		}
	}
	return tasks
}

func createTask(text string) db.Task {
	client := &http.Client{}
	task := db.Task{}
	task.Text = text
	task.Id = 0
	task.CreateTime = time.Now()
	JSONbytes, err := json.Marshal(task)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(
		"POST", "http://localhost:8000/tasks/create", bytes.NewReader(JSONbytes),
	)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	task = db.Task{} //initialized with zero value to check if unmarshalling works
	if resp.StatusCode == http.StatusOK {
		BodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(BodyBytes, &task)
		if err != nil {
			log.Fatal(err)
		}
	}
	return task
}

func completeTask(ID int) db.Task {
	client := &http.Client{}
	task := db.Task{}
	task.Id = ID
	task.CompleteTime.Time = time.Now()
	task.CompleteTime.Valid = true
	JSONbytes, err := json.Marshal(task)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(
		"PUT", "http://localhost:8000/tasks/complete/"+strconv.Itoa(task.Id), bytes.NewReader(JSONbytes),
	)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	completedTask := db.Task{}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(bodyBytes, &completedTask)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	return completedTask
}
