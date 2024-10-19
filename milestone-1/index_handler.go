package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Session")
	if err == nil {
		sessionID := cookie.Value
		if userData, ok := sessionsStore[sessionID]; ok {
			fmt.Fprintf(w, "Successfully authorized to access GitHub on your behalf: %s", userData.Login)
			return
		}
	}

	// Generate a random state string
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	// Set the state string in a cookie
	expiration := time.Now().Add(10 * time.Minute)
	stateCookie := http.Cookie{Name: "OAuthState", Value: state, Expires: expiration}
	http.SetCookie(w, &stateCookie)

	// Redirect to GitHub for authentication
	url := oauthConf.AuthCodeURL(state, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusFound)
}
