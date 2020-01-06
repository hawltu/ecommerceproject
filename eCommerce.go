package main

import (
	"html/template"
	"net/http"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("templates/*.html"))
}

func home(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "eCommerce.html", nil)
}
func about(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "about.html", nil)
}
func register(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "register.html", nil)
}
func login(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "login.html", nil)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets/css"))
	mux.Handle("/assets/css/", http.StripPrefix("/assets/css", fs))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/register", register)
	http.ListenAndServe("Localhost:8080", mux)
}
