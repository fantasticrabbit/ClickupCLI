package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	"reverse":        "reverse the sort order (bool)",
	"subtasks":       "output subtasks (bool)",
	"include-closed": "include closed tasks (bool)",
}
var taskIntFlags = map[string]string{
	"page":            "page numeber for lists of more than 100 tasks",
	"due_date_gt":     "filter for due date greater than supplied value",
	"due_date_lt":     "filter for due date less than supplied value",
	"date_created_gt": "filter for created date greater than supplied value",
	"date_created_lt": "filter for created date less than supplied value",
	"date_updated_gt": "filter for updated date greater than supplied value",
	"date_updated_lt": "filter for updated date less than supplied value",
}

var AllFlags = make(map[string]string)

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
		for flag := range AllFlags {
			flagtype := cmd.Flags().Lookup(flag).Value.Type()
			fmt.Println(flagtype) //debugging
			switch {
			case flagtype == "string":
				if x, _ := cmd.Flags().GetString(flag); x != "" {
					viper.BindPFlag(flag, cmd.Flags().Lookup(flag))
					fmt.Println("set", flagtype, "flag", flag) //debugging
				}
			case flagtype == "bool":
				if x, _ := cmd.Flags().GetBool(flag); x {
					viper.BindPFlag(flag, cmd.Flags().Lookup(flag))
					fmt.Println("set", flagtype, "flag", flag) //debugging
				}
			case flagtype == "int":
				if x, _ := cmd.Flags().GetInt(flag); x != 0 {
					viper.BindPFlag(flag, cmd.Flags().Lookup(flag))
					fmt.Println("set", flagtype, "flag", flag) //debugging
				}
			case flagtype == "stringSlice":
				if x, _ := cmd.Flags().GetStringSlice(flag); len(x) > 0 {
					viper.BindPFlag(flag, cmd.Flags().Lookup(flag))
					fmt.Println("set", flagtype, "flag", flag) //debugging
				}
			default:
				fmt.Println("did nothing", flag, flagtype)
			}

		}

		fmt.Println(tl.BuildPath()) //debugging
		internal.Request(tl)
	},
}

func init() {
	getCmd.AddCommand(tasksCmd)

	for flag, description := range taskStringFlags {
		tasksCmd.Flags().String(flag, "", description)
		AllFlags[flag] = description
	}
	for flag, description := range taskBoolFlags {
		tasksCmd.Flags().Bool(flag, false, description)
		AllFlags[flag] = description
	}
	for flag, description := range taskArrayFlags {
		tasksCmd.Flags().StringSlice(flag, nil, description)
		AllFlags[flag] = description
	}
	for flag, description := range taskIntFlags {
		tasksCmd.Flags().Int(flag, 0, description)
		AllFlags[flag] = description
	}
}
