package cmd

import (
	"fmt"
	"go_stuff/models"
	"go_stuff/storage"
	"time"

	"github.com/spf13/cobra"
)

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
		todos, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		var lastid int = 0
		for _, t := range todos {
			if t.ID > lastid {
				lastid = t.ID
			}
		}
		id := lastid + 1
		newTask := models.Task{ID: id, Description: taskDescrption, CreatedAt: time.Now(), Completed: false}
		todos = append(todos, newTask)
		err = storage.SaveTasks(todos)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
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
