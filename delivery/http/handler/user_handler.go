package handler

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/menu"

	//"github.com/amthesonofGod/Notice-Board/post"

	uuid "github.com/satori/go.uuid"
)

// UserHandler handles user requests
type UserHandler struct {
	tmpl    *template.Template
	userSrv menu.UserService
	//postSrv post.PostService
}

// NewUserHandler initializes and returns new NewUserHandler
func NewUserHandler(T *template.Template, US menu.UserService) *UserHandler {
	return &UserHandler{tmpl: T, userSrv: US}
}

// Index handle requests on /
func (uh *UserHandler) Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	uh.tmpl.ExecuteTemplate(w, "eCommerce.html", nil)

}

// Login handle requests on /login
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("session")

	if r.Method == http.MethodPost {

		username1 := r.FormValue("username")
		password := r.FormValue("password")

		users, _ := uh.userSrv.Users()

		for _, user := range users {
			if username1 == user.Username && password == user.Password {
				fmt.Println("authentication successfull! ")

				if err == http.ErrNoCookie {
					sID, _ := uuid.NewV4()
					cookie = &http.Cookie{
						Name:  "session",
						Value: sID.String(),
						Path:  "/",
					}
				}
				session := &entity.UserSession{}
				//session := &entity.UserSession{}
				session.UUID = cookie.Value
				session.UserID = user.Username

				/*_, errs := uh.userSrv.StoreSession(session)

				if len(errs) > 0 {
					panic(errs)
				}*/

				http.SetCookie(w, cookie)
				http.Redirect(w, r, "/home", http.StatusSeeOther)
				break

			} else {
				fmt.Println("No such user!")
			}
		}

		io.WriteString(w, cookie.String())

	} else {
		uh.tmpl.ExecuteTemplate(w, "index_signin_signup.html", nil)
	}
}

// CreateAccount handle requests on /signup-account
func (uh *UserHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("session")
	if r.Method == http.MethodPost {

		usr := &entity.User{}
		//usr.Username = r.FormValue("username")
		usr.Username = r.FormValue("username")
		usr.Mobile = r.FormValue("mobile")
		//
		usr.Shopname = r.FormValue("shopname")
		usr.Fname = r.FormValue("fullname")
		usr.Address = r.FormValue("address")
		usr.Email = r.FormValue("email")
		usr.Password = r.FormValue("password")
		// confirmpass := r.FormValue("confirmPassword")

		users, _ := uh.userSrv.Users()

		for _, user := range users {

			if usr.Username == user.Username {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				fmt.Println("This Email is already in use! ")
				return
			}
		}

	errs := uh.userSrv.StoreUser(usr)

		if len(errs) > 0 {
			panic(errs)
		}

		if err == http.ErrNoCookie {
			sID, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session",
				Value: sID.String(),
				Path:  "/",
			}
		}

		session := &entity.UserSession{}
		session.UUID = cookie.Value
		session.UserID = usr.Username

		/*_, errs = uh.userSrv.StoreSession(session)

		if len(errs) > 0 {
			panic(errs)
		}*/

		fmt.Println(usr)

		fmt.Println("User added to db")

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/home", http.StatusSeeOther)

	} else {
		uh.tmpl.ExecuteTemplate(w, "loggedin.html", nil)
	}

}

// Home handle requests on /home
/*func (uh *UserHandler) Home(w http.ResponseWriter, r *http.Request) {

	//get cookie
	_, err := r.Cookie("session")
	if err != nil {
		fmt.Println("no cookie")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	posts, _ := uh.postSrv.Posts()

	uh.tmpl.ExecuteTemplate(w, "", posts)
}
*/
// Logout Logs the user out
func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

	// get cookie
	cookie, err := r.Cookie("session")

	if err != http.ErrNoCookie {
		/*_, errs := uh.userSrv.DeleteSession(cookie.Value)
		// session.DeleteSession
		if len(errs) > 0 {
			panic(errs)
		}*/
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 302)
}
