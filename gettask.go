package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	CLICKUP_CLIENT_ID = os.Getenv("CLICKUP_CLIENT_ID")
)

func getClickUpTask(clickUpTaskID string, tokenValue string) {
	apiPath := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/", clickUpTaskID)

	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, apiPath, nil)

	req.Header.Add("Authorization", tokenValue)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", CLICKUP_CLIENT_ID)

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
