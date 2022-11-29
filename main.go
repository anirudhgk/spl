package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type plist struct {
	Id          int
	Name        string
	Company     string
	Category    string
	Description string
	Url         string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "use your password here"
	dbName := "epharmacy"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Aview(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM plist ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	list := plist{}
	res := []plist{}
	for selDB.Next() {
		var id int
		var name, company, category, description, url string
		err = selDB.Scan(&id, &name, &company, &category, &description, &url)
		if err != nil {
			panic(err.Error())
		}
		list.Id = id
		list.Name = name
		list.Company = company
		list.Category = category
		list.Description = description
		list.Url = url
		res = append(res, list)
	}
	tmpl.ExecuteTemplate(w, "Aview", res)
	defer db.Close()
}

func Uview(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM plist ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	list := plist{}
	res := []plist{}
	for selDB.Next() {
		var id int
		var name, company, category, description, url string
		err = selDB.Scan(&id, &name, &company, &category, &description, &url)
		if err != nil {
			panic(err.Error())
		}
		list.Id = id
		list.Name = name
		list.Company = company
		list.Category = category
		list.Description = description
		list.Url = url
		res = append(res, list)
	}
	tmpl.ExecuteTemplate(w, "Uview", res)
	defer db.Close()
}

func Ainsert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		company := r.FormValue("company")
		category := r.FormValue("category")
		description := r.FormValue("description")
		url := r.FormValue("url")
		insForm, err := db.Prepare("INSERT INTO plist(name, company, category, description, url) VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, company, category, description, url)
		log.Println("INSERT: Name: " + name + " | Company: " + company + "description, category, url")
	}
	defer db.Close()
	http.Redirect(w, r, "/aview", 301)
}

func Aupdate(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		company := r.FormValue("company")
		category := r.FormValue("category")
		description := r.FormValue("description")
		url := r.FormValue("url")
		id := r.FormValue("pid")
		insForm, err := db.Prepare("UPDATE plist SET name=?, company=?, category=?, description=?, url=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, company, category, description, url, id)
		log.Println("UPDATE: Name: " + name + " | Company: " + company)
	}
	defer db.Close()
	http.Redirect(w, r, "/aview", 301)
}

func Adelete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	list := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM plist WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(list)
	log.Println("DELETED")
	defer db.Close()
	http.Redirect(w, r, "/aview", 301)
}

func Ashow(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	pId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM plist WHERE id=?", pId)
	if err != nil {
		panic(err.Error())
	}
	list := plist{}
	for selDB.Next() {
		var id int
		var name, company, category, description, url string
		err = selDB.Scan(&id, &name, &company, &category, &description, &url)
		if err != nil {
			panic(err.Error())
		}
		list.Id = id
		list.Name = name
		list.Company = company
		list.Category = category
		list.Description = description
		list.Url = url
	}
	tmpl.ExecuteTemplate(w, "Ashow", list)
	defer db.Close()
}

func Ushow(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	pId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM plist WHERE id=?", pId)
	if err != nil {
		panic(err.Error())
	}
	list := plist{}
	for selDB.Next() {
		var id int
		var name, company, category, description, url string
		err = selDB.Scan(&id, &name, &company, &category, &description, &url)
		if err != nil {
			panic(err.Error())
		}
		list.Id = id
		list.Name = name
		list.Company = company
		list.Category = category
		list.Description = description
		list.Url = url
	}
	tmpl.ExecuteTemplate(w, "Ushow", list)
	defer db.Close()
}

func Aedit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	pId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM plist WHERE id=?", pId)
	if err != nil {
		panic(err.Error())
	}
	list := plist{}
	for selDB.Next() {
		var id int
		var name, company, category, description, url string
		err = selDB.Scan(&id, &name, &company, &category, &description, &url)
		if err != nil {
			panic(err.Error())
		}
		list.Id = id
		list.Name = name
		list.Company = company
		list.Category = category
		list.Description = description
		list.Url = url
	}
	tmpl.ExecuteTemplate(w, "Aedit", list)
	defer db.Close()
}

func Uadd(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	pId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM plist WHERE id=?", pId)
	if err != nil {
		panic(err.Error())
	}
	list := plist{}
	for selDB.Next() {
		var id int
		var name, company, category, description, url string
		err = selDB.Scan(&id, &name, &company, &category, &description, &url)
		if err != nil {
			panic(err.Error())
		}
		list.Id = id
		list.Name = name
		list.Company = company
		list.Category = category
		list.Description = description
		list.Url = url
		addForm, err := db.Prepare("INSERT INTO cart(id,name, company, category, description, url) VALUES(?,?,?,?,?,?)")
		addForm.Exec(id, name, company, category, description, url)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Name: " + name + " | Company: " + company + "description, category, url")
	}
	defer db.Close()
	http.Redirect(w, r, "/uview", 301)
}

func Ucart(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM cart ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	list := plist{}
	res := []plist{}
	for selDB.Next() {
		var id int
		var name, company, category, description, url string
		err = selDB.Scan(&id, &name, &company, &category, &description, &url)
		if err != nil {
			panic(err.Error())
		}
		list.Id = id
		list.Name = name
		list.Company = company
		list.Category = category
		list.Description = description
		list.Url = url
		res = append(res, list)
	}
	tmpl.ExecuteTemplate(w, "Ucart", res)
	defer db.Close()
}

func Udelete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	list := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM cart WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(list)
	log.Println("DELETED")
	defer db.Close()
	http.Redirect(w, r, "/ucart", 301)
}

func Anew(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Anew", nil)
}

func Aindex(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Aindex", nil)
}

func Uindex(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Uindex", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Home", nil)
}

func Alogin(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Alogin", nil)
}

func Ulogin(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Ulogin", nil)
}

func Admin(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	db := dbConn()
	sqlS := "SELECT * FROM admin WHERE username=? AND password=?"

	var id int
	row := db.QueryRow(sqlS, username, password)

	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		tmpl.ExecuteTemplate(w, "Alogin", nil)
	default:
		tmpl.ExecuteTemplate(w, "Aindex", nil)
	}

	defer db.Close()
}

func User(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	db := dbConn()
	sqlS := "SELECT * FROM user WHERE username=? AND password=?"

	var id int
	row := db.QueryRow(sqlS, username, password)

	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		tmpl.ExecuteTemplate(w, "Ulogin", nil)
	default:
		tmpl.ExecuteTemplate(w, "Uindex", nil)
	}
	defer db.Close()
}

func Signup(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		signForm, err := db.Prepare("INSERT INTO user(username, password) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		signForm.Exec(username, password)
		log.Println("INSERT: Name: " + username + " | pass: " + password + "description, category, url")
	}
	tmpl.ExecuteTemplate(w, "Ulogin", nil)
	defer db.Close()
}

func Usignup(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Usignup", nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "About", nil)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Contact", nil)
}

func Uupload(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Uupload", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		prescription := r.FormValue("prescription")
		uploadForm, err := db.Prepare("INSERT INTO upload(prescription) VALUES(?)")
		if err != nil {
			panic(err.Error())
		}
		uploadForm.Exec(prescription)
		log.Println("INSERT: prescription: " + prescription)
	}
	defer db.Close()
	http.Redirect(w, r, "/uindex", 301)
}

// all the routs are handled here
func main() {
	log.Println("Server started on: http://localhost:8083")
	http.HandleFunc("/", Home)
	http.HandleFunc("/alogin", Alogin)
	http.HandleFunc("/ulogin", Ulogin)
	http.HandleFunc("/admin", Admin)
	http.HandleFunc("/user", User)
	http.HandleFunc("/usignup", Usignup)
	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/about", About)
	http.HandleFunc("/contact", Contact)
	// admin routes
	http.HandleFunc("/aindex", Aindex)
	http.HandleFunc("/aview", Aview)
	http.HandleFunc("/anew", Anew)
	http.HandleFunc("/ainsert", Ainsert)
	http.HandleFunc("/ashow", Ashow)
	http.HandleFunc("/edit", Aedit)
	http.HandleFunc("/aupdate", Aupdate)
	http.HandleFunc("/adelete", Adelete)
	// customer routes
	http.HandleFunc("/uindex", Uindex)
	http.HandleFunc("/uview", Uview)
	http.HandleFunc("/ushow", Ushow)
	http.HandleFunc("/uupload", Uupload)
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/uadd", Uadd)
	http.HandleFunc("/ucart", Ucart)
	http.HandleFunc("/udelete", Udelete)
	http.ListenAndServe(":8083", nil)
}
