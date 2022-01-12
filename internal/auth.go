package internal

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/fantasticrabbit/ClickupCLI/utils"
	"github.com/pkg/browser"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

const (
	ProdAPIbase   = "https://app.clickup.com/api"
	ProdAPIbaseV2 = "https://app.clickup.com/api/v2"
)

// CheckToken returns True if a user auth token is availalble, otherwise false
func CheckToken() bool {
	if !viper.InConfig("token") || viper.GetString("token") == "" {
		return false
	} else {
		return true
	}
}

// SaveToken saves the API Access Token to the config file for re-use
func SaveToken(token string) {
	viper.Set("token", token)
	viper.WriteConfigAs(utils.GetConfigPath())
}

// GetToken retrieves client ID, client secret, and localhost port, and implements
// webserver to allow end-user to authenticate, returning authorization token
func GetToken() string {
	// Check for required config keys:
	if !(viper.IsSet("client_id")) {
		log.Fatalln("No Client ID provided, check configuration")
	}
	if !(viper.IsSet("client_secret")) {
		log.Fatalln("No Client Secret provided, check configuration")
	}

	ctx := context.Background()
	redirectURL := "http://localhost:" + viper.GetString("port")

	conf := &oauth2.Config{
		ClientID:     viper.GetString("client_id"),
		ClientSecret: viper.GetString("client_secret"),
		RedirectURL:  redirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  ProdAPIbase,
			TokenURL: ProdAPIbaseV2 + "/oauth/token",
		},
	}

	state := fmt.Sprint(rand.Int())
	messages := make(chan bool)
	authPath := conf.AuthCodeURL(state, oauth2.AccessTypeOnline)
	code := ""
	// callback handler, redirect from login is handled here
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// check that the state parameter matches
		if s, ok := r.URL.Query()["state"]; ok && s[0] == state {
			// auth code is received as query parameter
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
	if err := browser.OpenURL(authPath); err != nil {
		log.Fatalln("failed to open browser for authentication", err)
	}

	server := &http.Server{Addr: ":" + viper.GetString("port")}

	go func() {
		okToClose := <-messages
		if okToClose {
			if err := server.Shutdown(context.Background()); err != nil {
				log.Fatalln("Failed to shutdown server", err)
			}
		}
	}()

	log.Println(server.ListenAndServe())

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatalln(err)
	}

	return tok.AccessToken

}
