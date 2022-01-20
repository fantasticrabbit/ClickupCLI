package cmd

import (
	"strings"

	"github.com/fantasticrabbit/ClickupCLI/api"
	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var flessListsCmd = &cobra.Command{
	Use:   "folderless-lists SPACEID",
	Short: "get data for lists in a workspace not in a folder",
	Long: `Request JSON data for all folderless lists 
	by workspace ID`,
	Args: cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if authed := internal.CheckTokenExists(); !authed {
			internal.SaveToken(internal.GetToken())
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("archived", cmd.Flags().Lookup("archived"))
		l := api.ListRequest{
			SpaceID:  strings.Trim(args[0], " "),
			Archived: viper.GetBool("archived"),
		}
		api.Request(l)
	},
}

func init() {
	getCmd.AddCommand(flessListsCmd)
	flessListsCmd.Flags().BoolP("archived", "a", false, "include archived lists in output")
}
