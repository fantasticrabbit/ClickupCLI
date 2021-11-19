package cmd

import (
	"errors"
	"strings"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:   "tasks LIST_ID",
	Short: "get tasks from a list by supplying the list id",
	Long: `Request JSON data for a set of tasks in an 
	authorized Clickup workspace`,
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

		var tl = internal.TaskListRequest{
			ListID: strings.Trim(args[0], " "),
		}

		internal.Request(tl)
	},
}

func init() {
	getCmd.AddCommand(tasksCmd)

}
