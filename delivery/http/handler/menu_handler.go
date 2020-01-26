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

func (mh *MenuHandler) about(w http.ResponseWriter, r *http.Request) {
}
func (mh *MenuHandler) women(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "women.html", nil)
}
func (mh *MenuHandler) Register(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "register1.html", nil)
}
func (mh *MenuHandler) mennn(w http.ResponseWriter, r *http.Request) { 
	mh.tmpl.ExecuteTemplate(w, "men.html", nil)
}
func (mh *MenuHandler) Loging(w http.ResponseWriter, r *http.Request) {
		mh.tmpl.ExecuteTemplate(w, "login.html", nil)
}
func (mh *MenuHandler) tech(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "tech.html", nil)
}
func (mh *MenuHandler) upload(w http.ResponseWriter, r *http.Request){
	mh.tmpl.ExecuteTemplate(w, "upload.html", nil)
}
func (mh *MenuHandler)  kids(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "kids.html", nil)
}
func (mh *MenuHandler) home(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "eCommerce.html",nil)
}
/*func (mh *MenuHandler) upload(w http.ResponseWriter, r *http.Request){
	mh.tmpl.ExecuteTemplate(w, "upload.html", nil)
}*/
 