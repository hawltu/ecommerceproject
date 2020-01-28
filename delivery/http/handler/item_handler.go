package handler

import (
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"os"
	"path/filepath"
	"github.com/hawltu/project1/form"
	"github.com/hawltu/project1/user"
	"github.com/hawltu/project1/item"
	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/rtoken"
	"strconv"
	//"github.com/satori/go.uuid"
)

type ItemHandler struct {
	tmpl   *template.Template
	serSrv item.ItemService
	usrSer  user.UserService
	csrfSignKey    []byte
}

func NewItemrHandler(T *template.Template, US item.ItemService,USS user.UserService,csKey []byte) *ItemHandler {
	return &ItemHandler{tmpl: T, serSrv: US,usrSer: USS,csrfSignKey:csKey}
}
func (cph *ItemHandler) UploadItem(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("companypostsnew function invoked! ")
	cookie, err := r.Cookie("session")
	if r.Method == http.MethodPost {

		fmt.Println("post method verified! ")

		if err == nil {
			cookievalue := cookie.Value
			fmt.Println(cookievalue)
		}
	
		s, serr := cph.usrSer.Session(cookie.Value)

		if len(serr) > 0 {
			panic(serr)
		}

		fmt.Println("userid",s.UserID)
		post := &entity.Item{}
		post.UserID = s.UserID
		post.Name  = r.FormValue("name")
		post.Catagory = r.FormValue("catagory")
		xx := r.FormValue("catagory")
		fmt.Println("catagory",xx)
		post.Subcatagory = r.FormValue("subcatagory")
		x,_ := strconv.Atoi(r.FormValue("price"))
		post.Price = x
		y,_:= strconv.Atoi(r.FormValue("quantity"))
		post.Quantity = y
		fmt.Println("quantity is", y)
		//fmt.Println(post.Catagory)
		mf, fh, err := r.FormFile("image")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		post.Image = fh.Filename
		fmt.Println("the name of image",fh.Filename)
		writeFile1(&mf, fh.Filename)

		/*cmp, cerr := cph.usrSer.User(s.ID)

		if len(cerr) > 0 {
			fmt.Println("i am the error")
			panic(cerr)
		}*/

		//fmt.Println(cmp.Name)

	

		fmt.Println(post)
		_,errs := cph.serSrv.StoreItem(post)
		/*if len(errs) > 0{
			panic(errs)
		}*/
		
		///fmt.Println("post added to db")
		fmt.Println(errs)
		http.Redirect(w, r, "/log", http.StatusSeeOther)

	}else {
		fmt.Println("is else statement")
		cph.tmpl.ExecuteTemplate(w, "register.html", nil)

	}
}
func writeFile1(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "../../", "ui", "assets", "image", fname)
	image, err := os.Create(path)

	/*if err != nil {
		panic(err)
	}*/
	defer image.Close()
	io.Copy(image, *mf)
}


func (uh *ItemHandler) ItemByCatagoryMen(w http.ResponseWriter, r *http.Request){
	var mencategories  = []entity.Item{}
	itms,_ := uh.serSrv.Items()
	men := "men"
	for _, tt := range itms {
		if  tt.Catagory == men{
			mencategories = append(mencategories,tt)
			//uh.tmpl.ExecuteTemplate(w,"kids.html",kidcategories)
		}
	}
	uh.tmpl.ExecuteTemplate(w,"men.html",mencategories)
}
func (uh *ItemHandler) ItemByCatagoryKid(w http.ResponseWriter, r *http.Request) {
	var kidcategories   =  []entity.Item{}
	kid := "kid"
	itms,_ := uh.serSrv.Items()
	for _, tt := range itms{
		if  tt.Catagory == kid {
			kidcategories = append(kidcategories,tt)
		}
	}
	uh.tmpl.ExecuteTemplate(w,"kids.html",kidcategories)
}
func (uh *ItemHandler) ItemByCatagoryWomen(w http.ResponseWriter, r *http.Request) {
	itms,_ := uh.serSrv.Items()
	var womencategories   =  []entity.Item{}
	women := "women"
	for _, tt := range itms {
		if  tt.Catagory == women {
			womencategories = append(womencategories,tt)
			//uh.tmpl.ExecuteTemplate(w,"kids.html",kidcategories)
		}
	}
	uh.tmpl.ExecuteTemplate(w,"women.html",womencategories)
}
func (uh *ItemHandler) ItemByCatagoryTech(w http.ResponseWriter, r *http.Request) {
	var Techcategories   =  []entity.Item{}
	tech := "tech"
	itms,_ := uh.serSrv.Items()
	for _, tt := range itms {
		if  tt.Catagory == tech {
			Techcategories = append(Techcategories,tt)
			//uh.tmpl.ExecuteTemplate(w,"kids.html",kidcategories)
		}
	}
	uh.tmpl.ExecuteTemplate(w,"tech.html",Techcategories)
}
func (ach *ItemHandler) ItemUpdate(w http.ResponseWriter, r *http.Request) {
	//ach.tmpl.ExecuteTemplate(w,"update.html",nil)
	if r.Method == http.MethodGet {
		fmt.Println("here")
		idRaw := r.URL.Query().Get("id")
		fmt.Println("here is the problem")
		id, err := strconv.Atoi(idRaw)
		fmt.Println("id",id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
	   cat, errs := ach.serSrv.Item(id)
		if errs != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		
		   fmt.Println(cat)
		   ach.tmpl.ExecuteTemplate(w, "update.html", cat)

	}else if r.Method == http.MethodPost {

		pst := entity.Item{}

		//postid, _ := strconv.Atoi(r.FormValue("id"))

		//price, _ := strconv.Atoi(r.FormValue("price"))
		//quantity, _ := strconv.Atoi(r.FormValue("Quantity"))
		//pst.ID = uint(postid)
		//pst.Price = price
		//pst.Quantity = quantity
		pst.Name = r.FormValue("name")

		pst.Catagory = r.FormValue("catagory")
		pst.Subcatagory= r.FormValue("subcatagory")
		

		mf, fh, err := r.FormFile("image")

		if err != nil {
			panic(err)
		}

		defer mf.Close()

		pst.Image = fh.Filename

		writeFile1(&mf, pst.Image)

		errs := ach.serSrv.UpdateItem(pst)

		if err != nil {
			panic(errs)
		}

		http.Redirect(w, r, "/log", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "log", http.StatusSeeOther)
	}




}

func (it *ItemHandler) ItemDelete( w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		errs := it.serSrv.DeleteItem(id)
		if errs != nil{
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/log", http.StatusSeeOther)
}

func (it *ItemHandler) ItemBuyMen( w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		fmt.Println("here")
		idRaw := r.URL.Query()["id"][0]
		fmt.Println("id",idRaw)
		id,_:= strconv.Atoi(idRaw)
		fmt.Println("id",id);
		/*id,err:= strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}*/
		fmt.Println(id)
		itm,_ := it.serSrv.Item(id)
		it.tmpl.ExecuteTemplate(w,"checkout.html",itm)
	}
}
func (uh *ItemHandler) Logingg(w http.ResponseWriter, r *http.Request) {


	//get cookie
	/*s, err := r.Cookie("session")
	if err != nil {
		fmt.Println("no cookie")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	user, _ := uh.userSrv.Users()
*/
	uh.tmpl.ExecuteTemplate(w, "register.html", nil)
}
func (uh *ItemHandler) checking(w http.ResponseWriter, r *http.Request) {


	//get cookie
	/*s, err := r.Cookie("session")
	if err != nil {
		fmt.Println("no cookie")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	user, _ := uh.userSrv.Users()
*/
	uh.tmpl.ExecuteTemplate(w, "register.html", nil)
}
func (ach *ItemHandler) ItemUpdate1(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query()["id"][0]
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		cat, errs := ach.serSrv.Item(id)
		if errs != nil{
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		values := url.Values{}
		values.Add("id", idRaw)
		values.Add("name", cat.Name)
		//values.Add("price", cat.Price)
		values.Add("catagory", cat.Catagory)
		values.Add("subcatagory", cat.Subcatagory)
		values.Add("image", cat.Image)
		//values.Add("quantity",cat.Quantity)
		upCatForm := struct {
			Values   url.Values
			VErrors  form.ValidationErrors
			Item *entity.Item
			CSRF     string
		}{
			Values:   values,
			VErrors:  form.ValidationErrors{},
			Item:   &cat,
			CSRF:     token,
		}
		ach.tmpl.ExecuteTemplate(w, "update.html", upCatForm)
		return
	}
	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		
		// Validate the form contents
		updateCatForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		updateCatForm.Required("name","catagory","subcatagory","image")
		//updateCatForm.MinLength("catdesc", 10)
		updateCatForm.CSRF = token
		//price,_ := strconv.Atoi(r.FormValue("price"))
	    //quantity,_ := strconv.Atoi(r.FormValue("quantity"))
		catID, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		ctg := entity.Item{
			
			UserID:        uint(catID),
			Name:        r.FormValue("catname"),
			//Price: price,
			Image:       r.FormValue("image"),
			Catagory : r.FormValue("catagory"),
			Subcatagory : r.FormValue("subcatagory"),
			//Quantity : quantity,
		}
		mf, fh, err := r.FormFile("image")
		if err == nil {
			ctg.Image = fh.Filename
			writeFile1(&mf, ctg.Image)
		}
		if mf != nil {
			defer mf.Close()
		}
		errs := ach.serSrv.UpdateItem(ctg)
		if errs != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/log", http.StatusSeeOther)
		return
	}
}

func (uh *ItemHandler) Buy(w http.ResponseWriter, r *http.Request){
		fmt.Println("geting")
		idRaw := r.URL.Query()["id"][0]
		id, _ := strconv.Atoi(idRaw)

		itms,_ := uh.serSrv.Item(id)
		itms.Quantity = itms.Quantity - 1 
		fmt.Println("buying")
		uh.tmpl.ExecuteTemplate(w,"men.html",itms)
		//http.Redirect(w, r, "/log", http.StatusSeeOther)
	
}
