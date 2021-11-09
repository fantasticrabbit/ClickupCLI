package cmd

import (
	"log"
	"os"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	home, _ = os.UserHomeDir()
)

var rootCmd = &cobra.Command{
	Use:   "clickup",
	Short: "ClickupCLI allows access to ClickUp from the command line",
	Long: `ClickupCLI allows you to use data from Clickup to drive scripts,
	build tools, and send and receive data from your Clickup space.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clickup/config.yaml)")
}

// initConfig reads in config file, ENV variables if set, and determines authentication steps
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Set config in home/.clickup/config.yaml
		_, err := os.Stat(home + "/.clickup")
		if os.IsNotExist(err) {
			errDir := os.MkdirAll(home+"/.clickup", 0755)
			if errDir != nil {
				log.Fatalln("cannot create .clickup config folder:" + home + "/.clickup")
			}
		}
		viper.SetConfigFile(home + "/.clickup/config.yaml")
	}

	viper.SetEnvPrefix("CLICKUP")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		viper.ReadInConfig()
	}
	// Check for required config keys:
	if !(viper.IsSet("client_id")) {
		log.Fatalln("No Client ID provided, check configuration")
	}
	if !(viper.IsSet("client_secret")) {
		log.Fatalln("No Client Secret provided, check configuration")
	}
	viper.SetDefault("redirect_port", "4321")

	if !viper.InConfig("ctoken") || viper.GetString("ctoken") == "" {
		token, err := internal.GetToken(viper.GetString("client_id"), viper.GetString("client_secret"), viper.GetString("redirect_port"))
		if err != nil {
			log.Fatalln("auth failed")
		}
		viper.Set("cToken", token)
		viper.WriteConfigAs(home + "/.clickup/config.yaml")
	}
}
