/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go_stuff/storage"
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please provide the ID of the task to mark as done")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Enter an Int ID")
			return
		}
		todos, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Error Loading Data")
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
			fmt.Println("ID not Found")
			return
		}
		todos[index].Completed = true
		err = storage.SaveTasks(todos)
		if err != nil {
			fmt.Println("Error Saving Data")
			return
		}
		fmt.Println("Task Status Updated")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
