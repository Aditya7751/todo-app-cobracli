package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "to list down all tasks to to do",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"ID", "Description", "Created At"})
		data, err := os.ReadFile("data/todos.json")
		if err != nil {
			fmt.Println("Error while reading todos.json")
			return
		}
		var todos []Task
		err = json.Unmarshal(data, &todos)
		if err != nil {
			fmt.Println("Error While Unmarshalling Data")
			return
		}
		for _, t := range todos {
			table.Append([]string{strconv.Itoa(t.ID), t.Description, t.CreatedAt.Format("2006-01-02 15:04:05")})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
