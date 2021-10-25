package cmd

import (
	"fmt"
	"os"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	cfgFile string
	home, _ = os.UserHomeDir()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cucli",
	Short: "ClickupCLI allows access to ClickUp from the command line",
	Long: `ClickupCLI allows you to use data from Clickup to drive scripts,
	build tools, and send and receive data from your Clickup space.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("confirm auth status before calling")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/clickup/clickup.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home/.config directory with name "clickup.yaml" (without extension).
		viper.AddConfigPath(home + "/.config/clickup/")
		viper.SetConfigType("yaml")
		viper.SetConfigName("clickup")
	}

	viper.SetEnvPrefix("clickup")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		viper.ReadInConfig()
	} else {

		tok, err := internal.GetCUToken(viper.GetString("client_id"), viper.GetString("client_secret"), "4321")
		if err == nil {
			fmt.Println("auth succeeded")
		}
		viper.Set("cToken", tok)
		viper.WriteConfigAs(home + "/.config/clickup/clickup.yaml")
	}
}
