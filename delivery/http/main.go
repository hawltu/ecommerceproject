package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/hawltu/project1/entity"

	repository "github.com/hawltu/project1/user/repository"
	service "github.com/hawltu/project1/user/service"

	repositoryCamp "github.com/amthesonofGod/Notice-Board/company/repositoryCamp"
	serviceCamp "github.com/amthesonofGod/Notice-Board/company/serviceCamp"

	postRepos "github.com/hawltu/project1/item/repository"
	postServ "github.com/hawltu/project1/item/service"

	//appRepos "github.com/amthesonofGod/Notice-Board/application/repository"
	//appServ "github.com/amthesonofGod/Notice-Board/application/service"

	//reqRepos "github.com/amthesonofGod/Notice-Board/request/repository"
	//reqServ "github.com/amthesonofGod/Notice-Board/request/service"

	"github.com/hawltu/project1/delivery/http/handler"

	"github.com//hawltu/project1/rtoken"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


const (
	host     = "localhost"
	port     = 5180
	user     = "postgres"
	password = "hawltu"
	dbname   = "user1"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	
	dbconn, err := gorm.Open("postgres", psqlInfo)
	
	if err != nil {
	panic(err)
	}
	
	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseGlob("ui/templates/*"))

	/*userRepo := repository.NewUserRepositoryImpl(dbconn)
	userServ := service.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler()*/
	
    userSessionRepo := repository.NewSessionGormRepo(dbconn)
	userSessionsrv := service.NewSessionService(userSessionRepo)

	postRepo := postRepos.NewItemGormRepo(dbconn)
	postSrv := postServ.ItemServiceImpl(postRepo)

	userRepo := repository.NewUserGormRepo(dbconn)
	userSrv := service.NewUserService(userRepo)
	sess := configSess()

	userHandler := handler.NewUserHandler(tmpl, userSrv, postSrv, userSessionsrv, sess)

	/*adminCatgHandler := handler.NewUserHanlder(tmpl, userServ)
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
	http.HandleFunc("/upload",itemCatHandler.uploadedItem)*/
	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/",userHandler.Index)
	http.ListenAndServe(":8181", nil)
}
