package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of a Clickup workspace",
	Long:  `logout allows the user to delete the access token for accessing a Clickup workspace`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleted authentication token")
		viper.Set("ctoken", "")
		viper.WriteConfigAs(home + "/.clickup.yaml")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
