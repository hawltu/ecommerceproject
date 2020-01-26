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

	//repositoryCamp "github.com/amthesonofGod/Notice-Board/company/repositoryCamp"
	//serviceCamp "github.com/amthesonofGod/Notice-Board/company/serviceCamp"

	postRepos "github.com/hawltu/project1/item/repository"
	postServ "github.com/hawltu/project1/item/service"

	//appRepos "github.com/amthesonofGod/Notice-Board/application/repository"
	//appServ "github.com/amthesonofGod/Notice-Board/application/service"

	//reqRepos "github.com/amthesonofGod/Notice-Board/request/repository"
	//reqServ "github.com/amthesonofGod/Notice-Board/request/service"

	"github.com/hawltu/project1/delivery/http/handler"

	"github.com/hawltu/project1/rtoken"

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
func createTables(dbconn *gorm.DB) []error {

	// dbconn.DropTableIfExists(&entity.CompanySession{}, &entity.UserSession{})
	// errs := dbconn.CreateTable( &entity.Request{}, &entity.Application{}).GetErrors()
	errs := dbconn.CreateTable(&entity.Item{}, &entity.UserSession{}, &entity.User{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	
	dbconn, err := gorm.Open("postgres", psqlInfo)
	
	if err != nil {
	panic(err)
	}
	
	defer dbconn.Close()

	createTables(dbconn)

	/*if err := dbconn.Ping(); err != nil {
		panic(err)
	}*/

	tmpl := template.Must(template.ParseGlob("ui/templates/*"))

	/*userRepo := repository.NewUserRepositoryImpl(dbconn)
	userServ := service.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler()*/
	
    userSessionRepo := repository.NewSessionGormRepo(dbconn)
	userSessionsrv := service.NewSessionService(userSessionRepo)

	postRepo := postRepos.NewItemGormRepo(dbconn)
	postSrv := postServ.NewItemServiceImpl(postRepo)

	userRepo := repository.NewUserGormRepo(dbconn)
	userSrv := service.NewUserService(userRepo)
	sess := configSess()

	userHandler := handler.NewUserHandler(tmpl, userSrv, postSrv,userSessionsrv,sess)
	menuHandler := handler.NewMenuHandler(tmpl,userSrv)
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
	http.HandleFunc("/login.html",menuHandler.Loging)
	http.HandleFunc("/register.html",menuHandler.Register)
	http.HandleFunc("/login",userHandler.Login)
	http.HandleFunc("/log",userHandler.LoggedInn)
	http.HandleFunc("/register",userHandler.CreateAccount)
	http.ListenAndServe(":8181", nil)
}

func configSess() *entity.UserSession {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := rtoken.GenerateRandomID(32)
	signingString, err := rtoken.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &entity.UserSession{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}