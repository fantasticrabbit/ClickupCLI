package cmd

import (
	"errors"
	"strings"

	"github.com/fantasticrabbit/ClickupCLI/api"
	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var taskCmd = &cobra.Command{
	Use:   "task TASK_ID",
	Short: "get data for a single task by supplying it's task id",
	Long:  `Request JSON data for a single task in an authorized Clickup workspace`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("incorrect number of arguments")
		}
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		if authed := internal.CheckToken(); authed == false {
			internal.SaveToken(internal.GetToken())
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("custom", cmd.Flags().Lookup("custom"))
		viper.BindPFlag("subtasks", cmd.Flags().Lookup("subtasks"))

		var t = api.TaskRequest{
			TaskID:     strings.Trim(args[0], "#"),
			CustomTask: viper.GetBool("custom"),
			TeamID:     viper.GetString("team"),
			Subtasks:   viper.GetBool("subtasks"),
		}

		api.Request(t)
	},
}

func init() {
	getCmd.AddCommand(taskCmd)
	taskCmd.Flags().BoolP("custom", "c", false, "task id provided is a clickup custom task id")
	taskCmd.Flags().BoolP("subtasks", "s", false, "include subtasks in output")
}
