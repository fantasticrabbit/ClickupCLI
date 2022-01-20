package cmd

import (
	"strings"

	"github.com/fantasticrabbit/ClickupCLI/api"
	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list LISTID",
	Short: "get data for a list object by supplying it's list id",
	Long: `Request JSON data for a list objectin an authorized 
	Clickup workspace`,
	Args: cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if authed := internal.CheckToken(); !authed {
			internal.SaveToken(internal.GetToken())
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		l := api.ListRequest{
			ListID: strings.Trim(args[0], " "),
		}

		api.Request(l)
	},
}

func init() {
	getCmd.AddCommand(listCmd)
}
