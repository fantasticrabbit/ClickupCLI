package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Get data for a single task",
	Long:  `Request JSON data for a single task in an authorized Clickup workspace`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Incorrect arguments, usage: clickup get task 123456")
		}
		taskID := strings.Trim(args[0], "#")
		token := viper.GetString("ctoken")
		clientID := viper.GetString("client_id")
		fileFlag, _ := cmd.Flags().GetBool("file")
		data := internal.GetTask(taskID, token, clientID)
		filenm := "clickup_" + taskID + ".json"

		if !fileFlag {
			fmt.Println(string(data))
			return
		} else {

			err := os.WriteFile(filenm, data, 0644)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error writing task JSON")
			}
		}
	},
}

func init() {
	getCmd.AddCommand(taskCmd)

	taskCmd.Flags().BoolP("file", "f", false, "output to file clickup_<taskID>.json")
}
