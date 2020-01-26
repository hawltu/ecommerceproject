package handler

import (
	//"fmt"
	"html/template"
	//"io"
	"net/http"

	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/item"

	//"github.com/amthesonofGod/Notice-Board/post"

	//"github.com/satori/go.uuid"
)

type ItemHandler struct {
	tmpl    *template.Template
	userSrv item.ItemService
	
}

func NewItemrHandler(T *template.Template, US item.ItemService) *ItemHandler {
	return &ItemHandler{tmpl: T, userSrv: US}
}

