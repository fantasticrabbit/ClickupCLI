package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getClickUpTask(clickUpTaskID, tokenValue, clientID string) {
	apiPath := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/", clickUpTaskID)

	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, apiPath, nil)

	req.Header.Add("Authorization", tokenValue)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", clientID)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	//var resultMap map[string]interface{}
	//json.Unmarshal([]byte(resp_body), &resultMap)

	//fmt.Println(resultMap["name"])
	//fmt.Println(resultMap["custom_fields"])
	//fmt.Println(resp.Status)
	fmt.Print(string(resp_body))
}