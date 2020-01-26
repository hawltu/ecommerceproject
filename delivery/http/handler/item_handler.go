package handler

import (
	//"fmt"
	"html/template"
	//"io"
	//"github.com/hawltu/project1/ite
	"github.com/hawltu/project1/item"
	//"github.com/satori/go.uuid"
)

type ItemHandler struct {
	tmpl   *template.Template
	serSrv item.ItemService
}

func NewItemrHandler(T *template.Template, US item.ItemService) *ItemHandler {
	return &ItemHandler{tmpl: T, serSrv: US}
}
