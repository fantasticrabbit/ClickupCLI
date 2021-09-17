package main

import (
	"os"
)

const (
	argIndexTaskID = 1
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
