package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
		keys := []string{"team", "port"}
		for _, key := range keys {
			x, _ := cmd.Flags().GetString(key)
			if x != "" {
				viper.BindPFlag(key, cmd.Flags().Lookup(key))
				viper.WriteConfigAs(config_file)
				println("Saved", key, "in config file")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("team", "", "", "set the Team ID")
	setCmd.Flags().StringP("port", "", "", "set the Redirect URL Port number")
}
