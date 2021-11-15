package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

var (
	Client HTTPClient
)

//Requester interface needs an API path to request JSON data
type Requester interface {
	BuildPath() string
	GetJSON(string) []byte
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func init() {
	Client = &http.Client{}
}

//Gets JSON data for any struct that implements Requester interface
func getJSON(apiPath string) []byte {
	req, _ := http.NewRequest(http.MethodGet, apiPath, nil)
	req.Header.Add("Authorization", viper.GetString("token"))
	req.Header.Add("Content-Type", "application/json")
	resp, err := Client.Do(req)
	if err != nil {
		log.Fatalln("Errored when sending request to the server")
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	return resp_body
}

func FormatJSON(jsonresp []byte) string {
	var formattedJSON bytes.Buffer
	if err := json.Indent(&formattedJSON, jsonresp, "", "    "); err != nil {
		log.Fatalln(err)
	}
	return formattedJSON.String()
}
