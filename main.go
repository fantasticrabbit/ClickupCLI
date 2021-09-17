package main

import (
	"fmt"
	"os"
)

const (
	argIndexTaskID = 1
)

var (
	teamID = os.Getenv("CLICKUP_TEAM_ID")
)

func main() {
	if len(os.Args) < 2 {
		panic("Task id argument is missing")
	}

	taskID := os.Args[argIndexTaskID]

	//Authenticate user
	cToken := fetchUserToken()
	fmt.Println(string(cToken))
	//Get task as JSON
	getClickUpTask(taskID, cToken)
}
