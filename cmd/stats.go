package cmd

import (
	"fmt"

	"go_stuff/storage"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Shows task statistics and progress",
	Long:  `Shows completed and pending task statistics.`,
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		total := len(tasks)
		completed := 0

		for _, task := range tasks {
			if task.Completed {
				completed++
			}
		}

		pending := total - completed

		fmt.Printf("Total Tasks: %d\n", total)
		fmt.Printf("Completed: %d\n", completed)
		fmt.Printf("Pending: %d\n\n", pending)

		bar := progressbar.NewOptions(
			total,
			progressbar.OptionSetDescription("Progress"),
			progressbar.OptionShowCount(),
			progressbar.OptionSetWidth(30),
		)

		err = bar.Set(completed)
		if err != nil {
			fmt.Println("Error rendering progress bar:", err)
			return
		}

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
