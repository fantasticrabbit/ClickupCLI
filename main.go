package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	taskURL := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/", taskID)
	sToken := fetchUserToken()
	authString := fmt.Sprintf("&#34;%s&#34;", sToken)

	client := &http.Client{}

	req, _ := http.NewRequest("GET", taskURL, nil)

	req.Header.Add("Authorization", authString)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
