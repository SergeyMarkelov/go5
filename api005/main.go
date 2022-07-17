package main

import (
	"fmt"
	"go5/api005/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	err := models.ConnectDatabase()
	if err != nil {
		println("Can't connect to the database")
		return
	}

	r := mux.NewRouter()

	// API v5
	r.HandleFunc("/health-check", HealthCheck).Methods("GET")
	r.HandleFunc("/login", ShowLoginForm).Methods("GET")
	r.HandleFunc("/login", CheckPersonByLoginAndPassword).Methods("POST")
	r.HandleFunc("/hello", HelloForm).Methods("GET")
	r.HandleFunc("/no", NoHelloForm).Methods("GET")

	//r.HandleFunc("/registration", modelsRegistration).Methods("PUT")
	//r.HandleFunc("/getdata", models.GetPersons).Methods("PUT")

	http.ListenAndServe(":8080", r)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func CheckPersonByLoginAndPassword(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	Login := r.FormValue("Login")
	Password := r.FormValue("Password")

	person := models.CheckPersonByLoginAndPassword(Login, Password)
	if person.Login == "Isn't okay" {
		fmt.Fprintf(w, "")
		NoHelloForm(w, r)

	} else {
		fmt.Fprintf(w, "")
		HelloForm(w, r)
	}

}

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {

	type ViewData struct {
		Title   string
		Message string
	}

	data := ViewData{
		Title:   "Hello",
		Message: "=)",
	}
	tmpl, _ := template.ParseFiles("templates/login.html") // templates/login.html
	tmpl.Execute(w, data)
}

func HelloForm(w http.ResponseWriter, r *http.Request) {

	type ViewData struct {
		Title   string
		Message string
	}

	data := ViewData{
		Title:   "Hello",
		Message: "=)",
	}
	tmpl, _ := template.ParseFiles("templates/hello.html") // templates/hello.html
	tmpl.Execute(w, data)
}

func NoHelloForm(w http.ResponseWriter, r *http.Request) {
	type ViewData struct {
		Title   string
		Message string
	}
	data := ViewData{
		Title:   "",
		Message: "",
	}

	tmpl, _ := template.ParseFiles("templates/no.html") // templates/no.html
	tmpl.Execute(w, data)
}
