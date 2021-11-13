package cmd

import (
	"errors"
	"fmt"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "get data for a list of tasks by supplying it's list id",
	Long:  `Request JSON data for a list of tasks in an authorized Clickup workspace`,
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

		var l = internal.ListRequest{
			ListID:   string(args[0]),
			Subtasks: viper.GetBool("subtasks"),
		}
		fmt.Println(string(l.GetJSON(l.BuildPath())))
	},
}

func init() {
	getCmd.AddCommand(listCmd)
}
