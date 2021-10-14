package internal

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/pkg/browser"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func GetCUToken(clientID, clientSecret, localHostPort string) (string, error) {
	ctx := context.Background()
	redirectURL := "http://localhost:" + localHostPort
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://app.clickup.com/api",
			TokenURL: "https://app.clickup.com/api/v2/oauth/token",
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
		panic(fmt.Errorf("failed to open browser for authentication %s", err.Error()))
	}

	server := &http.Server{Addr: ":" + localHostPort}

	go func() {
		okToClose := <-messages
		if okToClose {
			if err := server.Shutdown(context.Background()); err != nil {
				log.Println("Failed to shutdown server", err)
			}
		}
	}()

	log.Println(server.ListenAndServe())

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile(home + "/.clickup/token.json")
	return tok.AccessToken, err

}
