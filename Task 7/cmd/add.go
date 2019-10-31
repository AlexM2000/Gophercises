package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"./dbProcessing/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		createdTask := createTask(task)
		fmt.Printf("Added \"%s\" to your task list.\n", createdTask.Text)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
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
