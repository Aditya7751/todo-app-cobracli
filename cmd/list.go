package cmd

import (
	"fmt"
	"go_stuff/storage"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	showCompleted  bool
	showPending    bool
	filterPriority string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `Lists tasks with optional filtering.`,
	Run: func(cmd *cobra.Command, args []string) {

		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{
			"ID",
			"Description",
			"Priority",
			"Created At",
			"Completed",
		})

		todos, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		for _, t := range todos {

			if showCompleted && !t.Completed {
				continue
			}

			if showPending && t.Completed {
				continue
			}

			if filterPriority != "" && t.Priority != filterPriority {
				continue
			}

			table.Append([]string{
				strconv.Itoa(t.ID),
				t.Description,
				t.Priority,
				t.CreatedAt.Format("2006-01-02 15:04:05"),
				strconv.FormatBool(t.Completed),
			})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(
		&showCompleted,
		"done",
		false,
		"show only completed tasks",
	)

	listCmd.Flags().BoolVar(
		&showPending,
		"pending",
		false,
		"show only pending tasks",
	)

	listCmd.Flags().StringVarP(
		&filterPriority,
		"priority",
		"p",
		"",
		"filter by priority (low, medium, high)",
	)
}
