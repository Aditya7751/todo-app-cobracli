package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds Task to TODO List",
	Long:  `Adds Task to TODO List.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a task description")
			return
		}
		taskDescrption := args[0]
		data, err := os.ReadFile("data/todos.json")
		if err != nil {
			fmt.Println("Error loading todos.json file")
			return
		}
		var todos []Task
		err = json.Unmarshal(data, &todos)
		if err != nil {
			fmt.Println("Error Unmarshalling todos.json file")
			return
		}
		var lastid int = 0
		for _, t := range todos {
			if t.ID > lastid {
				lastid = t.ID
			}
		}
		id := lastid + 1
		newTask := Task{ID: id, Description: taskDescrption, CreatedAt: time.Now()}
		todos = append(todos, newTask)
		jsonData, err := json.MarshalIndent(todos, "", "")
		if err != nil {
			fmt.Println("Error marshalling data:", err)
			return
		}
		err = os.WriteFile("data/todos.json", jsonData, 0666)
		if err != nil {
			fmt.Println("Error Writing Data")
			return
		}
		fmt.Println("Task Added Successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
