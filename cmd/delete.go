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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes tasks based on ID of the task",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Provide ID to delete")
			return
		}
		deleteID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Could Not Parse ID")
			return
		}
		var todos []Task
		data, err := os.ReadFile("data/todos.json")
		if err != nil {
			fmt.Println("Error Reading todos.json")
			return
		}
		err = json.Unmarshal(data, &todos)
		if err != nil {
			fmt.Println("Error Unmarshalling Data")
			return
		}
		var deleteIndex int = -1
		for i, t := range todos {
			if t.ID == int(deleteID) {
				deleteIndex = i
			}
		}
		if deleteIndex == -1 {
			fmt.Println("Index Not Found")
			return
		}
		todos = append(todos[:deleteIndex], todos[(deleteIndex+1):]...)
		jsonData, err := json.Marshal(todos)
		if err != nil {
			fmt.Println("Error Marshalling Data")
			return
		}
		err = os.WriteFile("data/todos.json", jsonData, 0666)
		if err != nil {
			fmt.Println("Error Writing Data")
			return
		}
		fmt.Println("Successfully Deleted Data")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
