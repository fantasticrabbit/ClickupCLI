package internal

import (
	"fmt"
)

type TaskListRequest struct {
	ListID          string
	Archived        bool     `cflag:"archived"`
	Page            int      `cflag:"page"`     // starts at 0, 100 per page
	Order_By        string   `cflag:"order-by"` // default is created, else: id, updated, due_date
	Reverse         bool     `cflag:"reverse"`
	Subtasks        bool     `cflag:"subtasks"`
	Statuses        []string `cflag:"statuses"`
	Include_closed  bool     `cflag:"closed"`
	Assignees       []string `cflag:"assignees"`       // need numberic ID's
	Due_date_gt     int      `cflag:"due-date-gt"`     // UNIX time in ms
	Due_date_lt     int      `cflag:"due-date-lt"`     // UNIX time in ms
	Date_created_gt int      `cflag:"date-created-gt"` // UNIX time in ms
	Date_created_lt int      `cflag:"date-created-lt"` // UNIX time in ms
	Date_updated_gt int      `cflag:"date-updated-gt"` // UNIX time in ms
	Date_updated_lt int      `cflag:"date-updated-lt"` // UNIX time in ms
	CustomFields    string   `cflag:"custom-fields"`
}

//BuildPath creates the API path for a task request
func (tl TaskListRequest) BuildPath() string {
	return fmt.Sprintf("%s/list/%s/task?", prodAPIbaseV2, tl.ListID)

}

//GetJSON accepts an API path and returns byte payload of JSON data
func (tl TaskListRequest) GetJSON(apiPath string) []byte {
	return getJSON(apiPath)
}
