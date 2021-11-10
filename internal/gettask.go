package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type TaskRequest struct {
	TaskID     string
	CustomTask bool
	TeamID     string
	Subtasks   bool
}

//Builds the API path for a Clickup task request
func (t TaskRequest) BuildPath() string {
	if !t.CustomTask {
		return fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/?include_subtasks=%t",
			t.TaskID, t.Subtasks)
	} else {
		return fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/?custom_task_ids=%t&team_id=%s&include_subtasks=%t",
			t.TaskID, t.CustomTask, t.TeamID, t.Subtasks)
	}
}

func (t TaskRequest) WriteOut(payload []byte) {
	if !viper.GetBool("file") {
		fmt.Println(string(payload))
	}
	err := os.WriteFile("clickup_"+t.TaskID+".json", payload, 0644)
	if err != nil {
		log.Fatalln("Error writing task JSON")
	}
}
