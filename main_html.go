package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var tpl *template.Template

// a struct for user validation
type userdetails struct {
	Username string
	Password string
}

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func readForm(r *http.Request) *userdetails {

	r.ParseForm()
	user := new(userdetails)
	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(user, r.PostForm)
	if decodeErr != nil {
		log.Print("error mapping parsed form data to a struct :", decodeErr)
	}

	return user
}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		err := tpl.ExecuteTemplate(w, "Login.html", nil)
		// err := parsedTemplate.Execute(w, nil)
		if err != nil {
			log.Fatal("cant execute html :", err)
			return
		}

	} else {
		user := readForm(r)
		fmt.Fprint(w, "Hello", user.Username, "!")
	}

}

func main() {

	// // using gorilla mux
	r := mux.NewRouter()

	r.HandleFunc("/login", login)

	// creating a file server for the static files
	fileServer := http.FileServer(http.Dir("static"))
	r.PathPrefix("/").Handler(http.StripPrefix("/static/", fileServer))

	// creating the server
	// router := mux.NewRouter()
	// router.HandleFunc("/", login).Methods("POST")
	// router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	http.ListenAndServe(":8080", r)
}
