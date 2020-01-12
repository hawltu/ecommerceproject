package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/hawltu/project1/delivery/http/handler"
	"github.com/hawltu/project1/menu/repository"
	"github.com/hawltu/project1/menu/service"
)

func main() {

	dbconn, err := sql.Open("postgres", "postgres://app_admin:P@$$w0rdD2@localhost/restaurantdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseGlob("ui/templates/*"))

	userRepo := repository.NewUserRepositoryImpl(dbconn)
	userServ := service.NewUserServiceImpl(userRepo)

	adminCatgHandler := handler.NewUserHanlder(tmpl, userServ)
	menuHandler := handler.NewMenuHandler(tmpl, userServ)

	itemRepo := repository.NewItemRepositoryImpl(dbconn)
	itemServ := service.NewItemServiceImpl(itemRepo)

	itemCatHandler := handler.NewItemHandler(tmpl,itemServ)
	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", menuHandler.Index)
	http.HandleFunc("/register.html",menuHandler.register)
	http.HandleFunc("/eCommerce.html",)
	http.HandleFunc("/about", menuHandler.about)
	//http.HandleFunc("/contact_us.html", menuHandler.contact)
	http.HandleFunc("/men", menuHandler.mennn)
	http.HandleFunc("/women",menuHandler.women)
	http.HandleFunc("/login.html",menuHandler.login)
	http.HandleFunc("/register", adminCatgHandler.CreateAccount)
	http.HandleFunc("/login",adminCatgHandler.Login)
	http.HandleFunc("/",adminCatgHandler.Index)
	http.HandleFunc("/upload.html",menuHandler.upload)
	http.HandleFunc("/upload",itemCatHandler.uploadedItem)
	http.ListenAndServe(":8181", nil)
}
