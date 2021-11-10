package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type Requester interface {
	BuildPath() string
}

//requests the JSON data for all fields in a Clickup task or subtask
func GetJSON(r Requester) []byte {
	apiPath := r.BuildPath()
	token := viper.GetString("ctoken")
	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, apiPath, nil)

	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp_body
}
