package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

var indexTemplate = `{{range $key,$value:=.Providers}}
    <p><a href="/auth/{{$value}}">Log in with {{index $.ProvidersMap $value}}</a></p>
{{end}}`

var userTemplate = `
	<p><a href="/logout/{{.Provider}}">logout</a></p>
	<p>Hello {{.Name}}</p>`

func handleProvider(w http.ResponseWriter, r *http.Request) {
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(w, gothUser)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func handleProviderCallback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	t, _ := template.New("foo").Parse(userTemplate)
	AddUserDetails(user, mux.Vars(r)["provider"])
	t.Execute(w, user)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}
	t, _ := template.New("foo").Parse(indexTemplate)
	t.Execute(w, providerIndex)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func handleAllUsers(w http.ResponseWriter, r *http.Request) {
	var userDetails []User
	var userID int
	var name, email, phone, createdOn, meta string
	users := GetAllUsers()
	for users.Next() {
		var user User
		err := users.Scan(&userID, &name, &email, &phone, &createdOn, &meta)
		if err != nil {
			log.Fatal(err)
		}
		user.UserID = userID
		user.Name = name
		user.Email = email
		user.Phone = phone
		user.CreatedOn = createdOn
		user.Meta = meta

		userDetails = append(userDetails, user)
	}
	json.NewEncoder(w).Encode(userDetails)
}
