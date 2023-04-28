package api

import (
	"net/http"
	"oath/api/handlers"
)

func SetUpAPI() {
	http.HandleFunc("/login/github/", handlers.GithubLoginHandler)
	http.HandleFunc("/login/github/callback", handlers.GithubCallbackHandler)
	http.HandleFunc("/loggedin", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoggedinHandler(w, r, "")
	})
}
