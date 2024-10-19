package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func githubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Verify the state parameter
	oauthStateCookie, err := r.Cookie("OAuthState")
	if err != nil || r.FormValue("state") != oauthStateCookie.Value {
		http.Error(w, "Invalid OAuth state", http.StatusBadRequest)
		return
	}

	// Exchange the code for an access token
	token, err := oauthConf.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		log.Printf("Token exchange error: %v", err)
		return
	}

	// Create a new GitHub client
	client := github.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.AccessToken},
	)))

	// Get the authenticated user's details
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		http.Error(w, "Failed to get user details", http.StatusInternalServerError)
		log.Printf("GitHub API error: %v", err)
		return
	}

	// Generate a session identifier
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		http.Error(w, "Failed to generate session ID", http.StatusInternalServerError)
		return
	}
	sessionID := base64.URLEncoding.EncodeToString(b)

	// Store user data in the session store
	sessionsStore[sessionID] = userData{
		Login:       *user.Login,
		accessToken: token.AccessToken,
	}

	// Set the Session cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "Session",
		Value:   sessionID,
		Expires: time.Now().Add(24 * time.Hour),
	})

	// Clear the OAuthState cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "OAuthState",
		Value:  "",
		MaxAge: -1,
	})

	// Redirect to the index page
	http.Redirect(w, r, "/", http.StatusFound)
}
