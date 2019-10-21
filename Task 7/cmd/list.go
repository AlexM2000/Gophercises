package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"

	"./dbProcessing/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all incomplete tasks.",
	Run: func(cmd *cobra.Command, args []string) {
			tasks := getTasks()
			for _, v := range tasks {
				fmt.Printf("%d. %s \n", v.Id, v.Text)
			}
		},
	}


func init() {
	RootCmd.AddCommand(listCmd)
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
