package cmd

import (
	"fmt"
	"os"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// versionCmd represents the version command
var gettaskCmd = &cobra.Command{
	Use:   "task",
	Short: "get a task",
	Long:  `Get a task in JSON format, based on providing a Clickup task ID`,
	Run: func(cmd *cobra.Command, args []string) {
		clientID := viper.GetString("CLIENT_ID")
		token := viper.GetString("cToken")
		taskID, _ := cmd.Flags().GetString("taskid")
		fileFlag, _ := cmd.Flags().GetBool("file")
		//if file flag=t, write file 'cutask_'+taskid+'.json'
		if !fileFlag {
			fmt.Println(string(internal.GetClickUpTask(taskID, token, clientID)))
			return
		} else {
			filenm := "clickup_" + taskID + ".json"
			data := internal.GetClickUpTask(taskID, token, clientID)
			err := os.WriteFile(filenm, data, 0644)
			if err != nil {
				fmt.Println("Error writing task JSON")
			}
		}
	},
}

func init() {
	getCmd.AddCommand(gettaskCmd)
	gettaskCmd.Flags().StringP("taskid", "t", "", "Clickup task ID to get")
	gettaskCmd.Flags().BoolP("file", "f", false, "output to file <taskID>.json")

}
