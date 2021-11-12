package cmd

import (
	"log"
	"os"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	home, _     = os.UserHomeDir()
	config_path = (home + "/.clickup")
	config_file = (home + "/.clickup/config.yaml")
)

var rootCmd = &cobra.Command{
	Use:   "clickup",
	Short: "ClickupCLI allows access to ClickUp from the command line",
	Long: `ClickupCLI allows you to use data from Clickup to drive scripts,
	build tools, and send and receive data from your Clickup space.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
	Version: "v0.1.7",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: "+config_file+")")
}

// initConfig reads in config file, ENV variables if set, and determines authentication steps
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		_, err := os.Stat(config_path)
		if os.IsNotExist(err) {
			errDir := os.MkdirAll(config_path, 0755)
			if errDir != nil {
				log.Fatalln("cannot create .clickup config folder:" + config_path)
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

func checkToken() {
	if !viper.InConfig("ctoken") || viper.GetString("ctoken") == "" {
		token, err := internal.GetToken()
		if err != nil {
			log.Fatalln("auth failed")
		}
		viper.Set("cToken", token)
		viper.WriteConfigAs(config_file)
	}
}
