package main

import (
	"fmt"
	"net/http"
)

type ClickupResponse struct {
	Task struct {
		Items []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			ID  string `json:"id"`
			URI string `json:"uri"`
		} `json:"items"`
	} `json:"tracks"`
}

func getClickUpTask(clickUpTaskID string, authHeader string) {
	bearerHeader := fmt.Sprintf("Bearer %s", cToken)
	apiPath := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/", clickUpTaskID)

	// Get JSON
	_, err = doJSONRequest(apiPath, http.MethodGet, nil, bearerHeader)
	if err != nil {
		return err
	}

}
