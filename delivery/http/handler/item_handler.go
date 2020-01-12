package handler

import (
	//"fmt"
	"html/template"
	//"io"
	"net/http"

	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/menu"

	//"github.com/amthesonofGod/Notice-Board/post"

	//"github.com/satori/go.uuid"
)

// handles the about items
type ItemHandler struct {
	tmpl    *template.Template
	userSrv menu.ItemService
	
}

func NewItemrHandler(T *template.Template, US menu.ItemService) *ItemHandler {
	return &ItemHandler{tmpl: T, userSrv: US}
}
func (uh *ItemHandler) uploadedItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	usr := &entity.Item{}
	//session := &entity.UserSession{}
	usr.image := r.FormValue("image")
	usr.catagory := r.FormValue("catagory")
	usr.subCatagory := r.FormValue("subCatagory")
	usr.price := r.FormValue("price")
	usr.quantity := r.FormValue("quanrity")
	usr.name := r.FormValue("name")


	errs := uh.userSrv.StoreItem(*usr)

		if len(errs) > 0 {
			panic(errs)
		}
	

}
