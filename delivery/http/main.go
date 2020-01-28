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
	"github.com/julienschmidt/httprouter"

	//appRepos "github.com/amthesonofGod/Notice-Board/application/repository"
	//appServ "github.com/amthesonofGod/Notice-Board/application/service"

	//reqRepos "github.com/amthesonofGod/Notice-Board/request/repository"
	//reqServ "github.com/amthesonofGod/Notice-Board/request/service"

	"github.com/hawltu/project1/delivery/http/handler"
	//"github.com/hawltu/project1/delivery/http/handler/api"
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
var tmpl *template.Template
func init() {
	tmpl = template.Must(template.ParseGlob("ui/templates/*"))
}

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

	




	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
    userSessionRepo := repository.NewSessionGormRepo(dbconn)
	userSessionsrv := service.NewSessionService(userSessionRepo)

	postRepo := postRepos.NewItemGormRepo(dbconn)
	postSrv := postServ.NewItemServiceImpl(postRepo)

	userRepo := repository.NewUserGormRepo(dbconn)
	userSrv := service.NewUserService(userRepo)
	sess := configSess()

	userHandler := handler.NewUserHandler(tmpl, userSrv, postSrv,userSessionsrv,sess,csrfSignKey)
	menuHandler := handler.NewMenuHandler(tmpl,userSrv)
	itemHandler := handler.NewItemrHandler(tmpl,postSrv,userSrv,csrfSignKey)
	cmtHandl :=    handler.NewUserHandler1(userSrv)
	cmtHand2 :=    handler.NewItemHandler1(postSrv)
	router := httprouter.New()

    router.GET("/v1/users", cmtHandl.GetUsers)
	router.GET("/v1/users/:id", cmtHandl.GetUser)
	router.PUT("/v1/users/:id", cmtHandl.PutUser)
	router.DELETE("/v1/users/:id", cmtHandl.DeleteUser)
	router.POST("/v1/users", cmtHandl.PostUser)

	router.GET("/v1/items", cmtHand2.GetItems)
	router.GET("/v1/items/:id", cmtHand2.GetItem)
	router.PUT("/v1/items/:id", cmtHand2.PutItem)
	router.POST("/v1/items", cmtHandl.PostUser)







	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/",userHandler.Index)
	http.HandleFunc("/login.html",menuHandler.Loging)
	http.HandleFunc("/register.html",menuHandler.Register)
	http.HandleFunc("/men.html",itemHandler.ItemByCatagoryMen)
	http.HandleFunc("/kids.html",itemHandler.ItemByCatagoryKid)
	http.HandleFunc("/women.html",itemHandler.ItemByCatagoryWomen)
	http.HandleFunc("/tech.html",itemHandler.ItemByCatagoryTech)
	http.HandleFunc("/login",userHandler.Loginn)
	http.HandleFunc("/upload.html",menuHandler.Upload)
	http.HandleFunc("/eCommerce.html",menuHandler.Home)
	http.HandleFunc("/upload",itemHandler.UploadItem)
	http.HandleFunc("/update.html",itemHandler.ItemUpdate)
	http.HandleFunc("/log",userHandler.LoggedInn)
	http.HandleFunc("/about.html",menuHandler.About)
	http.HandleFunc("/item/update",itemHandler.ItemUpdate)
	//http.HandleFunc("/item/update",menuHandler.Update)
	//http.HandleFunc()()
	http.HandleFunc("/registerr",itemHandler.Logingg)
	http.HandleFunc("/men/buy",itemHandler.ItemBuyMen)
	http.HandleFunc("/women/buy",itemHandler.ItemBuyMen)
	http.HandleFunc("/kids/buy",itemHandler.ItemBuyMen)
	http.HandleFunc("/tech/buy",itemHandler.ItemBuyMen)
	http.HandleFunc("/buy",itemHandler.Buy)
	//http.HandleFunc("/item/delete",itemHandler.Deleting)
	http.HandleFunc("/item/delete",itemHandler.ItemDelete)

	http.HandleFunc("/register",userHandler.CreateAccount)
	//http.ListenAndServe(":8181", router)
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