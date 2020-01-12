package main

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	"html/template"
	//"github.com/satori/go.uuid"
	_ "github.com/lib/pq"

	//net/http"
	"net/http"
	//elpers "../helpers"
	"errors"
	//repos "../repos"
	//  "github.com/gorilla/securecookie"
)

const (
	host     = "localhost"
	port     = 5180
	user     = "postgres"
	password = "hawltu"
	dbname   = "user1"
)

var tmplate = template.Must(template.ParseGlob("ui/templates/*"))

func register(w http.ResponseWriter, r *http.Request) {
	tmplate.ExecuteTemplate(w, "register1.html", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	tmplate.ExecuteTemplate(w, "eCommerce.html", nil)
}
func login(w http.ResponseWriter, r *http.Request) {
	tmplate.ExecuteTemplate(w, "login.html", nil)
}
func indexx(w http.ResponseWriter, r *http.Request) {

	tmplate.ExecuteTemplate(w, "register1.html", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	tmplate.ExecuteTemplate(w, "eCommerce.html", nil)
}
func contact(w http.ResponseWriter, r *http.Request) {
	tmplate.ExecuteTemplate(w, "contact_us.html", nil)
}
func about(w http.ResponseWriter, r *http.Request) {
	tmplate.ExecuteTemplate(w, "about.html", nil)
}
func upload(w http.ResponseWriter, r *http.Request) {
	tmplate.ExecuteTemplate(w, "upload.html", nil)
}
func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/register.html", indexx)
	mux.HandleFunc("/register", register1)
	mux.HandleFunc("/login.html", login)
	mux.HandleFunc("/login", logedin)
	mux.HandleFunc("/eCommerce.html", home)
	mux.HandleFunc("/upload.html", upload)
	mux.HandleFunc("/upload", uploadd)
	mux.HandleFunc("/contact_us.html", contact)
	mux.HandleFunc("/about.html", about)
	http.ListenAndServe(":8190", mux)
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
	//lastt := r.FormValue("username")
	uName := r.FormValue("username")
	email := r.FormValue("email")
	//fullname := r.FormValue("fulname")
	mobile_No := r.FormValue("Mobile")
	shopname := r.FormValue("shopname")
	address := r.FormValue("address")
	pwd := r.FormValue("password")
	tmplate.ExecuteTemplate(w, "eCommerce.html", nil)
	insertStatement := `INSERT INTO user1 (username,fullname, email, password,mobile,shopname,address) VALUES ( $1, $2, $3,$4,$5,$6,$7)`

	_, err = dbconn.Exec(insertStatement, &uName, &fname, &email, &pwd, &mobile_No, &shopname, &address)

	if err != nil {
		panic(err)
	}
	fmt.Println("fullname is: ", fname)
	fmt.Println("username is : ", uName)
	fmt.Println("email is : ", email)
	fmt.Println("mobile : ", mobile_No)
	fmt.Println("shopname : ", shopname)
	fmt.Println("address: ", address)
	fmt.Println("password: ", pwd)

}

type user1 struct {
	username string
	password string
	email    string
	mobile   string
	fullname string
	shopname string
	address  string
}

func logedin(w http.ResponseWriter, r *http.Request) {

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
	username1 := r.FormValue("username")
	pwd  := r.FormValue("password")

	userrr, err := users()
	if err != nil {
		panic(err)
	}

	for _, user := range userrr {
		if user.username == username1 && user.password ==  pwd {
			fmt.Println("suceed")
			tmplate.ExecuteTemplate(w,"logeedin.html",user)

		}

	}

}
func uploadd(w http.ResponseWriter, r *http.Request) {
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
	image := r.FormValue("image")
	//lastt := r.FormValue("username")
	catagory := r.FormValue("catagory")
	subCatagory := r.FormValue("subCatagory")
	//fullname := r.FormValue("fulname")
	price := r.FormValue("price")
	name := r.FormValue("name")

	tmplate.ExecuteTemplate(w, "eCommerce.html", nil)
	insertStatement := `INSERT INTO sellers (image,catagory, subCatagory, price,name) VALUES ( $1, $2, $3,$4,$5)`

	_, err = dbconn.Exec(insertStatement, &image, &catagory, &subCatagory, &price, &name)

	if err != nil {
		panic(err)
	}
}

func users() ([]user1, error) {
	//db query

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbconn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	rows, err := dbconn.Query("SELECT * FROM user1;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []user1{}

	for rows.Next() {
		category := user1{}
		err = rows.Scan(&category.username, &category.password, &category.fullname, &category.email, &category.mobile, &category.address, &category.shopname)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, category)
	}

	return ctgs, nil
}
