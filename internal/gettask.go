package internal

import (
	"fmt"
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
		return fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/", t.TaskID)
	} else {
		return fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/?custom_task_ids=true&team_id=%s", t.TaskID, t.TeamID)
	}
}
