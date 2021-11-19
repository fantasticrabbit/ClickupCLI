package internal

import (
	"fmt"
	"net/url"
)

type TaskListRequest struct {
	ListID          string
	Archived        bool
	Page            int
	Order_By        string
	Reverse         bool
	Subtasks        bool
	Statuses        []string
	Include_Closed  bool
	Assignees       []string
	Due_Date_gt     int
	Due_Date_lt     int
	Date_Created_gt int
	Date_Created_lt int
	Date_Updated_gt int
	Date_Updated_lt int
	CustomFields    string
}

//BuildPath creates the API path for a task request
func (tl TaskListRequest) BuildPath() string {
	noflags := true

	switch {
	case noflags:
		return fmt.Sprintf("%s/list/%s/task?", prodAPIbaseV2, tl.ListID)
	default:
		return url.PathEscape(fmt.Sprintf("%s/list/%s/task?"+
			"?archived=%t"+
			"&page=%d"+
			"&order_by=%s"+
			"&reverse=%t"+
			"&subtasks=%t"+
			"&statuses[]=%v"+
			"&include_closed=%t"+
			"&assignees[]=%v"+
			"&due_date_gt=%d"+
			"&due_date_lt=%d"+
			"&date_created_gt=%d"+
			"&date_created_lt=%d"+
			"&date_updated_gt=%d"+
			"&date_updated_lt=%d"+
			"&custom_fields[]=%v",
			prodAPIbaseV2,
			tl.ListID,
			tl.Archived,
			tl.Page,
			tl.Order_By,
			tl.Reverse,
			tl.Subtasks,
			tl.Statuses,
			tl.Include_Closed,
			tl.Assignees,
			tl.Due_Date_gt,
			tl.Due_Date_lt,
			tl.Date_Created_gt,
			tl.Date_Created_lt,
			tl.Date_Updated_gt,
			tl.Date_Updated_lt,
			tl.CustomFields))
	}
}

//GetJSON accepts an API path and returns byte payload of JSON data
func (tl TaskListRequest) GetJSON(apiPath string) []byte {
	return getJSON(apiPath)
}
