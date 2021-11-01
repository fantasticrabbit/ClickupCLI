package cmd

import (
	"fmt"
	"os"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Get data for a single task",
	Long:  `Request JSON data for a single task in an authorized Clickup workspace`,
	Run: func(cmd *cobra.Command, args []string) {
		clientID := viper.GetString("client_id")
		token := viper.GetString("ctoken")
		taskID, _ := cmd.Flags().GetString("taskid")
		fileFlag, _ := cmd.Flags().GetBool("file")

		if !fileFlag {
			fmt.Println(string(internal.GetTask(taskID, token, clientID)))
			return
		} else {
			filenm := "clickup_" + taskID + ".json"
			data := internal.GetTask(taskID, token, clientID)
			err := os.WriteFile(filenm, data, 0644)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error writing task JSON")
			}
		}
	},
}

func init() {
	getCmd.AddCommand(taskCmd)

	getCmd.Flags().StringP("taskid", "i", "", "Clickup task ID to get")
	getCmd.Flags().BoolP("file", "f", false, "output to file <taskID>.json")
}
