package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

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
		checkToken()
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("custom", cmd.Flags().Lookup("custom"))
		viper.BindPFlag("subtasks", cmd.Flags().Lookup("subtasks"))

		var t = internal.TaskRequest{
			TaskID:     strings.Trim(args[0], "#"),
			CustomTask: viper.GetBool("custom"),
			TeamID:     viper.GetString("team"),
			Subtasks:   viper.GetBool("subtasks"),
		}
		x, err := internal.FormatJSON(string(t.GetJSON(t.BuildPath())))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(x, err)
	},
}

func init() {
	getCmd.AddCommand(taskCmd)
	taskCmd.Flags().BoolP("custom", "c", false, "task id provided is a clickup custom task id")
	taskCmd.Flags().BoolP("subtasks", "s", false, "include subtasks in output")
}
