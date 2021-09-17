package main

import (
	"os"
)

const (
	argIndexTaskID = 1
)

var (
	teamID = os.Getenv("CLICKUP_TEAM_ID")
)

func main() {
	taskID := os.Args[argIndexTaskID]
	if taskID == "" {
		panic("Task id argument is missing")
	}

	//Authenticate user, get token
	cToken := fetchUserToken()

	//Get task as JSON
	getClickUpTask(taskID, cToken)
}
