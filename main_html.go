package main

import (
	//"fmt"
	"net/http"
	//"log"
	"html/template"
)

func login(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("static/login.html")
	parsedTemplate.Execute(w, nil)
}

func main() {

	// using http built-in router
	http.HandleFunc("/", login)

	// creating a file server for the static files

	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// creating the server

	s := &http.Server{
		Addr: ":8000",
	}

	s.ListenAndServe()
}
