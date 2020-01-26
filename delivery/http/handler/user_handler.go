package handler

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"io"
	"mime/multipart"
	"path/filepath"
	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/user"
	uuid "github.com/satori/go.uuid"

	"github.com/hawltu/project1/session"
	"github.com/hawltu/project1/item"
	// "github.com/amthesonofGod/Notice-Board/rtoken"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler handles user requests
type UserHandler struct {
	tmpl           *template.Template
	userSrv        user.UserService
	postSrv        item.ItemService
	sessionService user.SessionService
	userSess       *entity.UserSession
	loggedInUser   *entity.User
	//csrfSignKey    []byte
}

type contextKey string

var ctxUserSessionKey = contextKey("signed_in_user_session")

// NewUserHandler initializes and returns new NewUserHandler
func NewUserHandler(T *template.Template, US user.UserService, PS item.ItemService, sessServ user.SessionService, usrSess *entity.UserSession) *UserHandler {
	return &UserHandler{tmpl: T, userSrv: US, postSrv: PS, sessionService: sessServ, userSess: usrSess}
}
// Authenticated checks if a user is authenticated to access a given route
func (uh *UserHandler) Authenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ok := uh.loggedIn(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserSessionKey, uh.userSess)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// Index handle requests on /
func (uh *UserHandler) Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	uh.tmpl.ExecuteTemplate(w, "eCommerce.html", nil)

}

// func CheckPasswordHash(password, hash string) bool {
//     err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//     return err == nil
// }

// Login handle requests on /login
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	cookie, errc := r.Cookie("session")

	if r.Method == http.MethodPost {

		email := r.FormValue("username")
		password := r.FormValue("password")
		users, _ := uh.userSrv.Users()

		for _, user := range users {
			if email == user.UserName {
				if errc == bcrypt.ErrMismatchedHashAndPassword {
					fmt.Println("Your username or password is wrong")
					return
				}

				// match := CheckPasswordHash(password, user.Password)
				// fmt.Println("Match:   ", match)

				if errc == http.ErrNoCookie {
					sID, _ := uuid.NewV4()
					cookie = &http.Cookie{
						Name:  "session",
						Value: sID.String(),
						Path:  "/",
					}
				}

				session := &entity.UserSession{}
				session.UUID = cookie.Value
				session.UserID = user.ID

				_, errs := uh.userSrv.StoreSession(session)

				if len(errs) > 0 {
					panic(errs)
				}

				fmt.Println(user.Password)
				fmt.Println(password)

				fmt.Println("authentication successfull!")

				http.SetCookie(w, cookie)
				fmt.Println(cookie.Value)
				//uh.tmpl.ExecuteTemplate(w,"loggedin.html",nil)
				http.Redirect(w, r, "/log", http.StatusSeeOther)
				break
			} else {
				fmt.Println("user not found")
				// http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}

	} else {
		uh.tmpl.ExecuteTemplate(w, "eCommerce.html", nil)
	}
}

func (uh *UserHandler) loggedIn(r *http.Request) bool {
	if uh.userSess == nil {
		return false
	}
	userSess := uh.userSess
	c, err := r.Cookie(userSess.UUID)
	if err != nil {
		return false
	}
	ok, err := session.Valid(c.Value, userSess.SigningKey)
	if !ok || (err != nil) {
		return false
	}
	return true
}

// CreateAccount handle requests on /signup-account
func (uh *UserHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	cookie, errc := r.Cookie("session")

	if r.Method == http.MethodPost {

		usr := &entity.User{}
         
        usr.Mobile = r.FormValue("Mobile")
		usr.FName = r.FormValue("fname")
		usr.LName = r.FormValue("lname")
		usr.Email = r.FormValue("email")
		usr.UserName = r.FormValue("username")
		password := r.FormValue("password")
		usr.Shopname = r.FormValue("shopname")
		usr.Address   = r.FormValue("address")
		//usr.Image   = r.FormValue("image")
		//confirmpass := r.FormValue("confirmPassword")

		mf, fh, err := r.FormFile("image")

		if err != nil {
			panic(err)
		}

		defer mf.Close()

		usr.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		users, _ := uh.userSrv.Users()

		for _, user := range users {

			if usr.UserName == user.UserName {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				fmt.Println("This Email is already in use! ")
				return
			}
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			// singnUpForm.VErrors.Add("password", "Password Could not be stored")
			// uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			panic(err)
			return
		}

		usr.Password = string(hashedPassword)

		fmt.Println(usr.Password)
		_, errs := uh.userSrv.StoreUser(usr)

		if len(errs) > 0 {
			panic(errs)
		}

		if errc == http.ErrNoCookie {
			sID, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session",
				Value: sID.String(),
				Path:  "/",
			}
		}

		session := &entity.UserSession{}
		session.UUID = cookie.Value
		session.UserID = usr.ID

		_, errs = uh.userSrv.StoreSession(session)

		if len(errs) > 0 {
			panic(errs)
		}

		fmt.Println(usr)

		fmt.Println("User added to db")

		http.SetCookie(w, cookie)
		uh.tmpl.ExecuteTemplate(w, "login.html", nil)
		//http.Redirect(w, r, "/home", http.StatusSeeOther)

	} else {
		uh.tmpl.ExecuteTemplate(w, "eCommerce.html", nil)
	}

}

// Home handle requests on /home
func (uh *UserHandler) Home(w http.ResponseWriter, r *http.Request) {

	//get cookie
	_, err := r.Cookie("session")
	if err != nil {
		fmt.Println("no cookie")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//posts, _ := uh.postSrv.Posts()

	uh.tmpl.ExecuteTemplate(w, "eCommerce.html", nil)
}

// Logout hanldes the POST /logout requests
func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// userSess, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
	// session.Remove(userSess.UUID, w)
	// uh.sessionService.DeleteSession(userSess.UUID)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}


func writeFile(mf *multipart.File, fname string) {

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


func (uh *UserHandler) LoggedInn(w http.ResponseWriter, r *http.Request) {

	//get cookie
	_, err := r.Cookie("session")
	if err != nil {
		fmt.Println("no cookie")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//user, _ := uh.userSrv.Users()

	uh.tmpl.ExecuteTemplate(w, "register.html", nil)
}

//