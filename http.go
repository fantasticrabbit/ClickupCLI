package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client = http.Client{
	Timeout: time.Second * 2, // Timeout after 2 seconds
}

func doPostRequest(path string, values url.Values, authHeader string) (data []byte, err error) {
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, path, strings.NewReader(values.Encode()))
	if err != nil {
		return
	}

	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}
