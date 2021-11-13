package internal

import (
	"fmt"
	"net/url"
)

type ListRequest struct {
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
func (l ListRequest) BuildPath() string {
	noflags := true

	switch {
	case noflags:
		return fmt.Sprintf("%s/list/%s/", prodAPIbaseV2, l.ListID)
	default:
		return url.PathEscape(fmt.Sprintf("%s/list/%s/?task"+
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
			l.ListID,
			l.Archived,
			l.Page,
			l.Order_By,
			l.Reverse,
			l.Subtasks,
			l.Statuses,
			l.Include_Closed,
			l.Assignees,
			l.Due_Date_gt,
			l.Due_Date_lt,
			l.Date_Created_gt,
			l.Date_Created_lt,
			l.Date_Updated_gt,
			l.Date_Updated_lt,
			l.CustomFields))
	}
}

//GetJSON accepts an API path and returns byte payload of JSON data
func (t ListRequest) GetJSON(apiPath string) []byte {
	return getJSON(apiPath)
}
