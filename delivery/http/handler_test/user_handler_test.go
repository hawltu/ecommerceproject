/*package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
	"html/template"

	"github.com/hawltu/project1/delivery/http/handler"
	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/rtoken"
	//userRepo "github.com/hawltu/project1/user/repository"
	//userServ "github.com/hawltu/project1/user/service"
	postRepos "github.com/hawltu/project1/item/repository"
	postServ "github.com/hawltu/project1/item/service"
	"github.com/julienschmidt/httprouter"
	repository "github.com/hawltu/project1/user/repository"
	service "github.com/hawltu/project1/user/service"

)

func TestUsers(t *testing.T){
	tmpl := template.Must(template.ParseGlob("ui/templates/*"))
	userRepo := repository.NewMockUserRepo(nil)
	userSrv := service.NewUserService(userRepo)


	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
    userSessionRepo := repository.NewSessionGormRepo(nil)
	userSessionsrv := service.NewSessionService(userSessionRepo)

	postRepo := postRepos.NewItemGormRepo(nil)
	postSrv := postServ.NewItemServiceImpl(postRepo)
    
	sess := configSess()
	userHandler := handler.NewUserHandler(tmpl, userSrv, postSrv,userSessionsrv,sess,csrfSignKey)
	mux := httprouter.New()
	//mux.GET("/vl/user", userHandler.GetUsers)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/vl/user")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}
	var mockUsers []entity.User
	var users []entity.User
	_ = json.Unmarshal(body, &users)
	mockUsers = append(mockUsers, entity.UserMock)
	fmt.Println(mockUsers)
	fmt.Println(users)
	if !reflect.DeepEqual(mockUsers, users) {
		// t.Errorf("want body to contain \n%q, but\n%q", mockUsers, users)
		t.Errorf("not expected result")
	}
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
*/