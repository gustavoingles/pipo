package main

import "net/http"

func setupRoutes() {
	http.HandleFunc("POST /signup", signupHandler)
	http.HandleFunc("POST /login", loginHandler)
}