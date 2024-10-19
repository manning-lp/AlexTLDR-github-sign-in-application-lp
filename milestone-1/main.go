package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"log"
	"net/http"
	"os"
)

var oauthConf = &oauth2.Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),

	Scopes:      []string{"repo", "user"},
	Endpoint:    github.Endpoint,
	RedirectURL: "http://localhost:8080/github/callback",
}

func goDotEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {

	oauthConf.ClientID = goDotEnvVar("GITHUB_CLIENT_ID")
	oauthConf.ClientSecret = goDotEnvVar("GITHUB_CLIENT_SECRET")

	log.Printf("Loaded Client ID: %s", oauthConf.ClientID)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/github/callback", githubCallbackHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
