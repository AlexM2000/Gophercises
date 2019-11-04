package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"./dbProcessing/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete task with that number.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range args {
			id, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			completedTask := completeTask(id)
			fmt.Printf("Task &s has been completed\n", completedTask.Text)
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}

func completeTask(ID int) db.Task {
	client := &http.Client{}
	task := db.Task{
		ID:       ID,
		Complete: true,
	}
	JSONbytes, err := json.Marshal(task)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(
		"PUT", "http://localhost:8000/tasks/"+strconv.Itoa(task.ID), bytes.NewReader(JSONbytes),
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
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.StatusCode)
	}
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
	return completedTask
}
