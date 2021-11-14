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

//BuildPath creates the API path for a task request
func (t TaskRequest) BuildPath() string {
	if !t.CustomTask {
		return fmt.Sprintf("%s/task/%s/?include_subtasks=%t",
			prodAPIbaseV2, t.TaskID, t.Subtasks)
	} else {
		return fmt.Sprintf("%s/task/%s/?custom_task_ids=%t&team_id=%s&include_subtasks=%t",
			prodAPIbaseV2, t.TaskID, t.CustomTask, t.TeamID, t.Subtasks)
	}
}

//GetJSON accepts an API path and returns byte payload of JSON data
func (t TaskRequest) GetJSON(apiPath string) string {
	return getJSON(apiPath)
}
