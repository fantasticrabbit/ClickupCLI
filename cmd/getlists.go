package cmd

import (
	"errors"
	"strings"

	"github.com/fantasticrabbit/ClickupCLI/api"
	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listsCmd = &cobra.Command{
	Use:   "lists FOLDERID",
	Short: "get data for lists in a folder",
	Long:  `Request JSON data for all lists by folder ID`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires the folder-id argument")
		}
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		if authed := internal.CheckToken(); authed == false {
			internal.SaveToken(internal.GetToken())
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("archived", cmd.Flags().Lookup("archived"))
		l := api.ListRequest{
			FolderID: strings.Trim(args[0], " "),
			Archived: viper.GetBool("archived"),
		}
		api.Request(l)
	},
}

func init() {
	getCmd.AddCommand(listsCmd)
	listsCmd.Flags().BoolP("archived", "a", false, "include archived lists in output")
}
