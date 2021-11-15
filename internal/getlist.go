package internal

import (
	"fmt"
)

type ListRequest struct {
	ListID   string
	FolderID string
	SpaceID  string
	Archived bool
}

//BuildPath creates the API path for a task request
func (l ListRequest) BuildPath() string {
	switch {

	case l.FolderID != "":
		return fmt.Sprintf("%s/folder/%s/list?archived=%t", prodAPIbaseV2, l.FolderID, l.Archived)

	case l.SpaceID != "":
		return fmt.Sprintf("%s/space/%s/list?archived=%t", prodAPIbaseV2, l.SpaceID, l.Archived)

	default:
		return fmt.Sprintf("%s/list/%s/", prodAPIbaseV2, l.ListID)
	}
}

//GetJSON accepts an API path and returns byte payload of JSON data
func (t ListRequest) GetJSON(apiPath string) []byte {
	return getJSON(apiPath)
}
