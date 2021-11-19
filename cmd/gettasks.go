package cmd

import (
	"errors"
	"strings"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
)

//
var taskStringFlags = map[string]string{
	"order-by": "order by created, id, updated, or due_date",
}
var taskArrayFlags = map[string]string{
	"statuses":  "filter by an array of status values",
	"assignees": "filter tasks by specified assignee IDs",
}
var taskBoolFlags = map[string]string{
	"archived":       "display active or archived tasks (bool)",
	"reversed":       "reverse the sort order (bool)",
	"subtasks":       "output subtasks (bool)",
	"include_closed": "include closed tasks (bool)",
}
var taskIntFlags = map[string]string{
	"due_date_gt":     "filter for due date greater than supplied value",
	"due_date_lt":     "filter for due date less than supplied value",
	"date_created_gt": "filter for created date greater than supplied value",
	"date_created_lt": "filter for created date less than supplied value",
	"date_updated_gt": "filter for updated date greater than supplied value",
	"date_updated_lt": "filter for updated date less than supplied value",
}

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
		//use reflect to get the flag type, use switch case to get flag and add to struct? or just

		internal.Request(tl)
	},
}

func init() {
	getCmd.AddCommand(tasksCmd)
	for flag, description := range taskStringFlags {
		tasksCmd.Flags().StringP(flag, "", "", description)
	}
	for flag, description := range taskBoolFlags {
		tasksCmd.Flags().BoolP(flag, "", false, description)
	}
	for flag, description := range taskArrayFlags {
		tasksCmd.Flags().StringArrayP(flag, "", nil, description)
	}

	for flag, description := range taskIntFlags {
		tasksCmd.Flags().IntP(flag, "", 0, description)
	}

}
