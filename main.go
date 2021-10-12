package main

import (
	"flag"
	"os"

	"github.com/fantasticrabbit/authcu"
)

var (
	appID      = os.Getenv("CLICKUP_CLIENT_ID")
	appSecret  = os.Getenv("CLICKUP_CLIENT_SECRET")
	redURLport = "4321"
	taskID     string
	fileout    string
)

func init() {
	const (
		taskUsage = "the Clickup Task ID to fetch"
		fileUsage = "the filename to store JSON output, otherwise StdOut"
	)

	flag.StringVar(&taskID, "task", "", taskUsage)
	flag.StringVar(&taskID, "t", "", taskUsage+" (shorthand)")
	flag.StringVar(&fileout, "file", "", fileUsage)
	flag.StringVar(&fileout, "f", "", fileUsage+" (shorthand)")
}

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
	cToken, _ := authcu.GetCUToken(appID, appSecret, redURLport)

	//Get task as JSON
	getClickUpTask(taskID, cToken, appID)
}
