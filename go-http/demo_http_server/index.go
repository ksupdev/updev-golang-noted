package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// http://localhost:8090/
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Home, %q", html.EscapeString(r.URL.Path))
	})

	// http://localhost:8090/profile
	http.HandleFunc("/profile", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Profile, %q", html.EscapeString(r.URL.Path))
	})

	//http://localhost:8090/login?username=admin&password=password
	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "login, %s  %s ", r.URL.Query().Get("username"), r.URL.Query().Get("password"))
	})

	log.Fatal(http.ListenAndServe(":8090", nil))
}
