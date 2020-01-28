package handler

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"io"
	"mime/multipart"
	"path/filepath"
	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/user"
	uuid "github.com/satori/go.uuid"

	"github.com/hawltu/project1/session"
	
	"github.com/hawltu/project1/form"
	
	"github.com/hawltu/project1/rtoken"
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
	csrfSignKey    []byte
}

type contextKey string

var ctxUserSessionKey = contextKey("signed_in_user_session")

// NewUserHandler initializes and returns new NewUserHandler
func NewUserHandler(T *template.Template, US user.UserService, PS item.ItemService, sessServ user.SessionService, usrSess *entity.UserSession,csKey []byte) *UserHandler {
	return &UserHandler{tmpl: T, userSrv: US, postSrv: PS, sessionService: sessServ, userSess: usrSess,csrfSignKey: csKey}
}

func (ach *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	categories, errs := ach.userSrv.Users()
	if errs != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
	     Users      []entity.User
		CSRF       string
	}{
		Values:     nil,
		VErrors:    nil,
		Users: categories,
		CSRF:       token,
	}
	ach.tmpl.ExecuteTemplate(w, "register.html", tmplData)
}

func CheckPasswordHash(password, hash string) bool {
     err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
     return err == nil
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

// Login handle requests on /login
/*func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	cookie, errc := r.Cookie("session")

	if r.Method == http.MethodPost {

		email := r.FormValue("username")
		password := r.FormValue("password")
		users, _ := uh.userSrv.Users()
		fmt.Println("username from form",email)
		for _, user := range users {
			/*if email != user.UserName{
				var x  = "INVALID USERNAME OR PASSWORD"
				uh.tmpl.ExecuteTemplate(w,"login.html",x)
				return
			}*/
			//err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.FormValue("password")))
			/*if email == user.UserName  {
				err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.FormValue("password")))
				if err == bcrypt.ErrMismatchedHashAndPassword  {
					fmt.Println("Your username or password is wrong")
					var x  = "INVALID USERNAME OR PASSWORD"
					uh.tmpl.ExecuteTemplate(w,"login.html",x)
					return
				}
					
				
				match := CheckPasswordHash(password, user.Password)
				fmt.Println("Match:   ", match)

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
				//break
				break
			}
			
		}

	} else {
		uh.tmpl.ExecuteTemplate(w, "eCommerce.html", nil)
	}
}
*/
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
	uh.tmpl.ExecuteTemplate(w, "eCommerce.html", nil)
}

/*func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	userSess, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
	session.Remove(userSess.UUID, w)
	uh.sessionService.DeleteSession(userSess.UUID)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}*/

// Logout hanldes the POST /logout requests
/*func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// userSess, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
	// session.Remove(userSess.UUID, w)
	// uh.sessionService.DeleteSession(userSess.UUID)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}*/


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
	s, err := r.Cookie("session")
	if err != nil {
		fmt.Println("no cookie")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	s1, _ := uh.userSrv.Session(s.Value)
	itms := []entity.Item{}
	user1,_ := uh.postSrv.Items()
	for _,tt := range user1{
		if tt.UserID == s1.UserID {
			itms = append(itms,tt)
		}
	}
	uh.tmpl.ExecuteTemplate(w, "register.html",itms)
}
func (uh *UserHandler) Loginn(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	cookie, errc := r.Cookie("session")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		loginForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "login.html", loginForm)
		return
	}
	if r.Method == http.MethodPost {
		// Parse the form data
		/*err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}*/
		//email := r.FormValue("username")
		//fmt.Println("username from form",email)
		loginForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		
		user, errs := uh.userSrv.UserByUserName(r.FormValue("username"))
		
		
		if len(errs) > 0 {
			loginForm.VErrors.Add("generic", "Your username address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "login.html", loginForm)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.FormValue("password")))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			loginForm.VErrors.Add("generic", "Your username address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "login.html", loginForm)
			return
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
		session.UserID = user.ID

		_, errs = uh.userSrv.StoreSession(session)
		fmt.Println("i am here")
		if len(errs) > 0 {
			panic(errs)
		}

		//fmt.Println(user.Password)
		//fmt.Println(password)

		fmt.Println("authentication successfull!")

		http.SetCookie(w, cookie)
		fmt.Println(cookie.Value)
		///uh.tmpl.ExecuteTemplate(w,"loggedin.html",nil)
		http.Redirect(w, r, "/log", http.StatusSeeOther)


		/*uh.loggedInUser = usr
		claims := rtoken.Claims(usr.Email, uh.userSess.Expires)
		session.Create(claims, uh.userSess.UUID, uh.userSess.SigningKey, w)
		newSess, errs := uh.sessionService.StoreSession(uh.userSess)
		if len(errs) > 0 {
			loginForm.VErrors.Add("generic", "Failed to store session")
			uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}
		uh.userSess = newSess
		roles, _ := uh.userService.UserRoles(usr)
		if uh.checkAdmin(roles) {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)*/
	}
}

