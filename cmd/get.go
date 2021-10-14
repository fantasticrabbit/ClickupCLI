package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets JSON from Clickup",
	Long:  `get retrieves data in JSON format from the Clickup API`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().String("task", "", "get a task based on Clickup task ID")
	getCmd.Flags().String("output", "", "name of file to output JSON")

}
