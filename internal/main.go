package internal

import (
	"flag"
	"os"
)

var (
	appID      = os.Getenv("CLICKUP_CLIENT_ID")
	appSecret  = os.Getenv("CLICKUP_CLIENT_SECRET")
	redURLport = "4321"
	taskID     string
	fileout    string
)

func main() {
	flag.Parse()

	if taskID == "" {
		panic("Task id argument is missing")
	}
	if appID == "" {
		panic("Missing Clickup Client ID env variable, server not setup properly")
	}
	if appSecret == "" {
		panic("Missing Clickup Client Secret env variable, server not setup properly")
	}

	//Authenticate user, get token
	cToken, _ := GetCUToken(appID, appSecret, redURLport)

	//Get task as JSON
	getClickUpTask(taskID, cToken, appID)
}
