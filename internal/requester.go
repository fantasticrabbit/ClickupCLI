package internal

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

//Requester interface needs an API path to request JSON data
type Requester interface {
	BuildPath() string
	GetJSON(string) []byte
	WriteOut([]byte)
}

//Gets JSON data for any struct that implements Requester interface
func getJSON(apiPath string) []byte {
	token := viper.GetString("ctoken")
	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, apiPath, nil)

	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Errored when sending request to the server")
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp_body
	//	fileFlag := viper.GetBool("file")
	//	if !fileFlag {
	//		fmt.Println(string(resp_body))
	//	}
}

func (t TaskRequest) GetJSON(apiPath string) []byte {
	return getJSON(apiPath)
}
