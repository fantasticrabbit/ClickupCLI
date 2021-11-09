package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var taskCmd = &cobra.Command{
	Use:   "task TASK_ID [-f]",
	Short: "get data for a single task by supplying it's task id",
	Long:  `Request JSON data for a single task in an authorized Clickup workspace`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("incorrect number of arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		taskID := strings.Trim(args[0], "#")
		token := viper.GetString("ctoken")
		clientID := viper.GetString("client_id")
		fileFlag, _ := cmd.Flags().GetBool("file")
		data := internal.GetTask(taskID, token, clientID)
		filenm := "clickup_" + taskID + ".json"

		if !fileFlag {
			fmt.Println(string(data))
			return
		} else {
			err := os.WriteFile(filenm, data, 0644)
			if err != nil {
				log.Fatalln("Error writing task JSON")
			}
		}
	},
}

func init() {
	getCmd.AddCommand(taskCmd)
	taskCmd.Flags().BoolP("file", "f", false, "output to file clickup_<taskID>.json")
}
