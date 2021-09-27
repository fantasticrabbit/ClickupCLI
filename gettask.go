package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Jeffail/gabs"
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
	data, _ := ioutil.ReadAll(resp.Body)

	jsonParsed, err := gabs.ParseJSON(data)
	if err != nil {
		panic(err)
	}

	// Search JSON
	fmt.Println("Get Project Name:\t", jsonParsed.Path("name").Data())
	chainVal := jsonParsed.Path("custom_fields.Chains.value").Data()
	fmt.Println(chainVal)
	// Iterate over config options to find chainVal
	// for chainVal==jsonParsed.Path("custom_fields.Chains.type_config.options")

	//fmt.Println("Get Chain Name:\t", jsonParsed.Path("name").Data())
	//fmt.Println("Get Project Name:\t", jsonParsed.Path("name").Data())
	//fmt.Println("Get Project Name:\t", jsonParsed.Path("name").Data())
	//fmt.Println("Get Project Name:\t", jsonParsed.Path("name").Data())

	//fmt.Println("Get value of Country:\t", jsonParsed.Search("employees", "address", "country").Data())
	//fmt.Println("ID of first employee:\t", jsonParsed.Path("employees.employee.0.id").String())
	//fmt.Println("Check Country Exists:\t", jsonParsed.Exists("employees", "address", "countryCode"))

	//fmt.Print(string(data))
}
