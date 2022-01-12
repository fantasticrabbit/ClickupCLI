package cmd

import (
	"errors"

	"github.com/fantasticrabbit/ClickupCLI/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// add new config options here with the flag command and help
// message for the setting. only string flags allowed
var configOptions = map[string]string{
	"team":  "set the Team ID",
	"port":  "set the Redirect URL Port number",
	"token": "set the Auth Token manually",
}

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "sets config options",
	Long: `set is used to configure extended options and save
	them to the config file`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("incorrect number of arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for flag := range configOptions {
			value, _ := cmd.Flags().GetString(flag)
			if value != "" {
				viper.BindPFlag(flag, cmd.Flags().Lookup(flag))
				viper.WriteConfigAs(utils.GetConfigFile())
				println("Saved", flag, "in config file")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	for flag, description := range configOptions {
		setCmd.Flags().StringP(flag, "", "", description)
	}
}
