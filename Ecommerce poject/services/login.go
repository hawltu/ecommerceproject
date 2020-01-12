
func loginhand(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	userlog = r.FormValue("username")  // Data from the form
	pwdlog = r.FormValue("password") // Data from the form

	

	dbuser := `SELECT username FROM user1`        // DB simulation
	dbpass := `SELECT password FROM user1` 	     // DB simulation

	_, err = dbconn.Exec(dbuser,  &userlog)
	_, err = dbconn.Exec(dbuser,  &pwdlog)
	if err != nil {
		panic(err)
	}
	if userlog == dbuser && pwdlog == dbpass {
		fmt.Fprintln(w, "Login succesful!")
		http.Redirect(w, r, "/logedin.html", http.StatusSeeOther)
		return
	} else {
		fmt.Fprintln(w, "Login failed!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}