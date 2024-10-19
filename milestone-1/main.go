package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var sessionsStore = make(map[string]userData)

var oauthConf = &oauth2.Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	Scopes:       []string{"repo", "user"},
	Endpoint:     github.Endpoint,
	RedirectURL:  "http://localhost:8080/github/callback",
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/github/callback", githubCallbackHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
