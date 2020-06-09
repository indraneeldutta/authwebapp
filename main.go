package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/linkedin"
	"github.com/markbates/goth/providers/twitter"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	key := "something"   // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"),
		linkedin.New(os.Getenv("LINKEDIN_KEY"), os.Getenv("LINKEDIN_SECRET"), "http://localhost:3000/auth/linkedin/callback"),
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),
	)

	initDb()
}

var (
	m = map[string]string{
		"github":   "Github",
		"linkedin": "Linkedin",
		"twitter":  "Twitter",
	}
)

// ProviderIndex struct
type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleHome).Methods("GET")
	router.HandleFunc("/auth/{provider}", handleProvider).Methods("GET")
	router.HandleFunc("/auth/{provider}/callback", handleProviderCallback).Methods("GET")
	router.HandleFunc("/logout/{provider}", handleLogout).Methods("GET")
	router.HandleFunc("/user/all", handleAllUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
