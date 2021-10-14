package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var gettaskCmd = &cobra.Command{
	Use:   "task",
	Short: "get a task",
	Long:  `Get a task in JSON format, based on providing a Clickup task ID`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
		main.getClickUpTask(taskID, cToken, appID)
	},
}

func init() {
	getCmd.AddCommand(gettaskCmd)

	gettaskCmd.Flags().StringP("taskid", "t", "", "Clickup task ID to get")
	gettaskCmd.Flags().BoolP("file", "f", false, "output to file <taskID>.json")

	main.getClickUpTask()
}
