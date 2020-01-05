package main

import (
	"database/sql"
	"fmt"
	"html/template"

	_ "github.com/lib/pq"

		"net/http"

)

const (
	host     = "localhost"
	port     = 8181
	user     = "postgres"
	password = "password"
	dbname   = "user1"
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
func index(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "register.html", nil)
}
func login(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "login.html", nil)
}

func register1(w http.ResponseWriter, r *http.Request) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbconn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fname := r.FormValue("fullname")
	uName := r.FormValue("username")
	email := r.FormValue("email")
	mobile_No := r.FormValue("Mobile")
	shopname := r.FormValue("shopname")
	address := r.FormValue("address")
	pwd := r.FormValue("password")
	templ.Execute(w, nil)
	insertStatement := `INSERT INTO user1 (username,fullname, email, password,mobile,shopname,address) VALUES ( $1, $2, $3,$4,$5,$6,$7)`
	//insert1 := 'INSERT INTO user1 (username,fullname, email, password,mobile,shopname,address) VALUES ( $1, $2, $3,$4,$5,$6,$7)'
	//_, err = dbconn.Exec(insertStatement)
	_, err = dbconn.Exec(insertStatement, &uName, &fname, &email, &pwd, &mobile_No, &shopname, &address)

	if err != nil {
		panic(err)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	templ.Execute(w, nil)
	uName := r.FormValue("username")
	email := r.FormValue("email")
	fullname := r.FormValue("fullname")
	mobile_No := r.FormValue("Mobile")
	shopname := r.FormValue("shopname")
	address := r.FormValue("address")
	pwd := r.FormValue("password")

	fmt.Println(w, "Username for Register : ", uName)
	fmt.Println(w, "Email for Register : ", email)
	fmt.Println(w, "address for Register : ", address)
	fmt.Println(w, "shopname for Register : ", shopname)
	fmt.Println("mobileNo", mobile_No)
	fmt.Println("fullname", fullname)
	fmt.Println("pwd", pwd)
	
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets/css"))
	mux.Handle("/assets/css/", http.StripPrefix("/assets/css", fs))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/register.html", index)
	mux.HandleFunc("/register", register1)
	mux.HandleFunc("/login", login)
	http.ListenAndServe("Localhost:8080", mux)
}
