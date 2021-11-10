package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

//Requester interface needs an API path to request JSON data
type Requester interface {
	BuildPath() string
	WriteOut([]byte)
}

//Gets JSON data for any struct that implements Requester interface
func GetJSON(r Requester) {
	apiPath := r.BuildPath()
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

	fileFlag := viper.GetBool("file")
	if !fileFlag {
		fmt.Println(string(resp_body))
	} else {
		r.WriteOut(resp_body)
	}
}
