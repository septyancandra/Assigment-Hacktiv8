package main

import "net/http"

const USERNAME = "batman"
const PASSWORD = "secret"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("something wrong"))
		return false
	}
	isValid := (username == USERNAME && password == PASSWORD)
	if !isValid {
		w.Write([]byte("something wrong username/password"))
		return false
	}
	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is ALlowed"))
		return false
	}
	return true
}
