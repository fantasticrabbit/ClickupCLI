package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTask(clickUpTaskID, tokenValue, clientID string) []byte {
	apiPath := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/", clickUpTaskID)

	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, apiPath, nil)

	req.Header.Add("Authorization", tokenValue)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", clientID)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp_body
}
