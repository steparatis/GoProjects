package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "login.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file: ", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, "Log in")
	if err != nil {
		fmt.Println("Error when executing template: ", err)
		return
	}
}

var users = map[string]string{
	"przemek": "password",
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if users[email] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Login successful")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Login failed")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)
	default:
		w.Write([]byte("Hello world"))
	}
}

func main() {
	http.HandleFunc("/", handler)
	//http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	http.ListenAndServe("", nil)
}
