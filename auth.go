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

	//check for env vars and die
	if clientID == "" && clientSecret == "" {
		panic(fmt.Errorf("environment vars for Clickup client ID and client secret are missing"))
	}

	code := ""                                                         // initialize authorization code - received in callback
	state := fmt.Sprint(rand.Int())                                    // local state parameter for cross-site request forgery prevention
	path := fmt.Sprintf(clickupLoginURL, clientID, redirectURL, state) // loginURL
	messages := make(chan bool)                                        // channel for signaling that server shutdown can be done

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
		// redirect user's browser to Clickup home page
		http.Redirect(w, r, "https://app.clickup.com/", http.StatusSeeOther)
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
	params.Add("client_id", clientID)
	params.Add("client_secret", clientSecret)
	params.Add("code", code)

	data, err := doPostRequest(
		"https://app.clickup.com/api/v2/oauth/token",
		params,
		authHeader,
	)

	if err == nil {
		response := AuthResponse{}
		if err = json.Unmarshal(data, &response); err == nil {
			// token parsed successfully
			return response.AccessToken
		}
	}
	panic(fmt.Errorf("unable to acquire Clickup user token"))
}
