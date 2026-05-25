/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "to update task description",
	Long:  `Use in the format update id newDescription`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Enter ID and New Description")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Enter a Number for ID")
		}
		var newDescription string = args[1]
		var found bool = false
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
		var index int = -1
		for i, t := range todos {
			if t.ID == id {
				index = i
				found = true
				break
			}
		}
		if found == false || index == -1 {
			fmt.Println("ID not Found")
			return
		}
		todos[index].Description = newDescription
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
		fmt.Println("Task Updated Successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
