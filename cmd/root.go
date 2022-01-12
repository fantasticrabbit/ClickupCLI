package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "clickup",
	Short: "ClickupCLI allows access to ClickUp from the command line",
	Long: `ClickupCLI allows you to use data from Clickup to drive scripts,
	build tools, and send and receive data from your Clickup space.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "specify a config file")
}

// initConfig reads in config file and available ENV variables
func initConfig() {
	var (
		home, _     = os.UserHomeDir()
		config_path = filepath.Join(home, ".clickup")
		config_file = filepath.Join(home, ".clickup", "config.yaml")
	)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		_, err := os.Stat(config_path)
		if os.IsNotExist(err) {
			err := os.MkdirAll(config_path, 0755)
			if err != nil {
				log.Fatalf("cannot create %s: %v", config_path, err)
			}
		}
		viper.SetConfigFile(config_file)
	}

	viper.SetEnvPrefix("CLICKUP")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		viper.ReadInConfig()
	}

	viper.SetDefault("port", "4321")
}
