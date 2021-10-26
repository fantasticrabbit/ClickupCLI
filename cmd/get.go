package cmd

import (
	"fmt"
	"os"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets JSON from Clickup",
	Long:  `get retrieves data in JSON format from the Clickup API`,
	Run: func(cmd *cobra.Command, args []string) {
		clientID := viper.GetString("client_id")
		token := viper.GetString("ctoken")
		taskID, _ := cmd.Flags().GetString("task")
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
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("task", "t", "", "Clickup task ID to get")
	getCmd.Flags().BoolP("file", "f", false, "output to file <taskID>.json")
}
