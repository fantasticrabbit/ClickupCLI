package main

import (
	"os"

	"github.com/fantasticrabbit/authcu"
)

const (
	argIndexTaskID = 1
)

var (
	appID      = os.Getenv("CLICKUP_CLIENT_ID")
	appSecret  = os.Getenv("CLICKUP_CLIENT_SECRET")
	redURLport = "4321"
)

func main() {
	taskID := os.Args[argIndexTaskID]
	if taskID == "" {
		panic("Task id argument is missing")
	}

	//Authenticate user, get token
	cToken, _ := authcu.GetCUToken(appID, appSecret, redURLport)

	//Get task as JSON
	getClickUpTask(taskID, cToken, appID)
}
