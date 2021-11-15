package cmd

import (
	"errors"
	"fmt"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list {LISTID | --folder=FOLDERID | --space=SPACEID} [-a]",
	Short: "get data for a list object(s) by supplying it's list id",
	Long: `Request JSON data for a list object or folders of list objects 
	in an authorized Clickup workspace`,
	Args: func(cmd *cobra.Command, args []string) error {
		switch {
		case len(args) == 0:
			if viper.GetString("folder") != "" && viper.GetString("space") != "" {
				return errors.New("do not provide both folder and space flags")
			}
			return nil

		case len(args) == 1:
			if viper.GetString("folder")+viper.GetString("space") != "" {
				return errors.New("do not provide folder or space flags if outputting a specific list")
			}
			return nil

		default:
			return errors.New("incorrect number of arguments")

		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		checkToken()
	},
	Run: func(cmd *cobra.Command, args []string) {

		l := internal.ListRequest{
			FolderID: viper.GetString("folder"),
			SpaceID:  viper.GetString("space"),
			Archived: viper.GetBool("archived"),
		}
		if len(args) == 1 {
			l.ListID = args[0]
		}

		fmt.Println(internal.FormatJSON(l.GetJSON(l.BuildPath())))
	},
}

func init() {
	getCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("archived", "a", false, "include archived lists in output")
	listCmd.Flags().StringP("folder", "", "", "get data on a group of lists by folder")
	listCmd.Flags().StringP("space", "", "", "get data on a group of lists by workspace")
	viper.BindPFlag("archived", listCmd.Flags().Lookup("archived"))
	viper.BindPFlag("folder", listCmd.Flags().Lookup("folder"))
	viper.BindPFlag("space", listCmd.Flags().Lookup("space"))
}
