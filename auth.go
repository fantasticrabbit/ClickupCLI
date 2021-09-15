package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"

	"github.com/pkg/browser"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

var (
	clientID     = os.Getenv("CLICKUP_CLIENT_ID")
	clientSecret = os.Getenv("CLICKUP_CLIENT_SECRET")
	authHeader   = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))
)

func fetchUserToken() string {
	const (
		redirectURL     = "http://localhost:4321"
		clickupLoginURL = "https://app.clickup.com/api?client_id=%s&redirect_uri=%s&state=%s"
	)

	if clientID == "" && clientSecret == "" {
		panic(fmt.Errorf("Clickup client ID and secret missing"))
	}

	// authorization code - received in callback
	code := ""

	// local state parameter for cross-site request forgery prevention
	state := fmt.Sprint(rand.Int())

	// loginURL
	path := fmt.Sprintf(clickupLoginURL, clientID, redirectURL, state)

	// channel for signaling that server shutdown can be done
	messages := make(chan bool)

	// callback handler, redirect from login is handled here
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// check that the state parameter matches
		if s, ok := r.URL.Query()["state"]; ok && s[0] == state {
			// code is received as query parameter
			if codes, ok := r.URL.Query()["code"]; ok && len(codes) == 1 {
				// save code and signal shutdown
				code = codes[0]
				messages <- true
			}
		}

		// redirect user's browser to spotify home page
		// http.Redirect(w, r, "https://app.clickup.com/", http.StatusSeeOther)
	})

	// open user's browser to login page
	if err := browser.OpenURL(path); err != nil {
		panic(fmt.Errorf("failed to open browser for authentication %s", err.Error()))
	}

	server := &http.Server{Addr: ":4321"}
	// go routine for shutting down the server
	go func() {
		okToClose := <-messages
		if okToClose {
			if err := server.Shutdown(context.Background()); err != nil {
				log.Println("Failed to shutdown server", err)
			}
		}
	}()
	// start listening for callback - we don't continue until server is shut down
	log.Println(server.ListenAndServe())

	// authentication complete - fetch the access token
	params := url.Values{}
	//params.Add("grant_type", "authorization_code")
	params.Add("code", code)
	params.Add("redirect_uri", redirectURL)
	data, err := doPostRequest(
		"https://app.clickup.com/api/v2/oauth/token",
		params,
		authHeader,
	)
	if err == nil {
		response := AuthResponse{}
		if err = json.Unmarshal(data, &response); err == nil {
			// happy end: token parsed successfully
			return response.AccessToken
		}
	}
	panic(fmt.Errorf("unable to acquire Clickup user token"))
}
