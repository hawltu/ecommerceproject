package handler

import (
	"html/template"
	"net/http"

	"github.com/hawltu/project1/user"
)

type MenuHandler struct {
	tmpl        *template.Template
	categorySrv user.UserService
}

func NewMenuHandler(t *template.Template,cat user.UserService) *MenuHandler  {
	return &MenuHandler{tmpl:t,categorySrv : cat}
}

func (mh *MenuHandler) Index(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "eCommerce.html", nil)
}

func (mh *MenuHandler) About(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "about.html", nil)
}

func (mh *MenuHandler) Register(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "register1.html", nil)
}
func (mh *MenuHandler) Men(w http.ResponseWriter, r *http.Request) { 
	mh.tmpl.ExecuteTemplate(w, "men.html", nil)
}
func (mh *MenuHandler) Loging(w http.ResponseWriter, r *http.Request) {
		mh.tmpl.ExecuteTemplate(w, "login.html", nil)
}
func (mh *MenuHandler) Upload(w http.ResponseWriter, r *http.Request){
	mh.tmpl.ExecuteTemplate(w, "upload.html", nil)
}
func (mh *MenuHandler) Home(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "eCommerce.html",nil)
}
func (mh *MenuHandler) Update(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "update.html",nil)
}
