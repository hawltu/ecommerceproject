func upload(w http.ResponseWriter, r *http.Request) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbconn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	iname := r.FormValue("named")
	iprice := r.FormValue("price")
	icategory := r.FormValue("category")
	isubcategory := r.FormValue("subcategory")
	
	templ.Execute(w, nil)
	insertStatement := `INSERT INTO items (name,price, category, subcategory) VALUES ( $n, $p, $c,$s)`
	_, err = dbconn.Exec(insertStatement, &iName, &iprice, &icategory, &isubcategory)

	if err != nil {
		panic(err)
	}
	

}
