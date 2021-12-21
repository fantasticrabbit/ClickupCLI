package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets JSON from Clickup",
	Long:  `get retrieves data in JSON format from the Clickup API`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
