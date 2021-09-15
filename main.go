package main

import (
	"fmt"
	"os"
)

const (
	argIndexTaskID = 1
)

type Track struct {
	Artist string
	Title  string
}

func main() {
	if len(os.Args) < 1 {
		panic("Task id argument is missing")
	}

	taskID := os.Args[argIndexTaskID]
	sToken := fetchUserToken()

	getTaskID(taskID, sToken)

	fmt.Printf("https://app.clickup.com/t/%s\n", taskID)
}
