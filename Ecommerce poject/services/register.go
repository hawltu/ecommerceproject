package project1

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)
var templ := template.Must(template.ParseFiles("register.html"));

func registerHandler(w http.ResponseWriter, r *http.Request) {
	templ3.Execute(w, nil)
	r.ParseForm()
	uName := r.FormValue("uername")
	email := r.FormValue("email")
	fullname := r.FormValue("fulname")
	mobile_No := r.FormValue("Moble")
	shopname := r.FormValue("shopname")
	address := r.FormValue("address")
	pwd := r.FormValue("password")
	//r.FormValue("address") = uName;
	connString := "dbname=<postgress> sslmode=disable"
	db, err := sql.Open("postgres", connString)
	sqlStatement := `
	INSERT INTO user1 ( password)
	VALUES ($1,$@)`
	id := 0
	err = db.QueryRow(sqlStatement, 244).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("hyyyyyy")

	fmt.Println(w, "Username for Register : ", uName)
	fmt.Println(w, "Email for Register : ", email)
	fmt.Println(w, "address for Register : ", address)
	fmt.Println(w, "shopname for Register : ", shopname)
	fmt.Println("mobileNo", mobile_No)
	fmt.Println("fullname", fullname)
	fmt.Println("pwd", pwd)
	fmt.Println(w, "This fields can not be blank!")
}
