package cmd

import (
	"fmt"
	"go_stuff/storage"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	description string
	priority    string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task fields",
	Long:  `Update task description or priority using flags.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Enter task ID")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Enter a valid numeric ID")
			return
		}

		if description == "" && priority == "" {
			fmt.Println("Provide at least one flag: --description or --priority")
			return
		}

		todos, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		index := -1

		for i, t := range todos {
			if t.ID == id {
				index = i
				break
			}
		}

		if index == -1 {
			fmt.Println("Task ID not found")
			return
		}

		if description != "" {
			todos[index].Description = description
		}

		if priority != "" {
			todos[index].Priority = priority
		}

		err = storage.SaveTasks(todos)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}

		fmt.Println("Task updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(
		&description,
		"description",
		"d",
		"",
		"new task description",
	)

	updateCmd.Flags().StringVarP(
		&priority,
		"priority",
		"p",
		"",
		"task priority",
	)
}
