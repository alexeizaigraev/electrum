// electrum project main.go
package main

import (
	"database/sql"
	"electrum/control_otbor"
	"electrum/db"
	"electrum/people"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Db *sql.DB

func main() {

	Db, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		//return err
		panic(err)
	}
	defer Db.Close()

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/menu_people", people.PagePeopleMenu)
	router.HandleFunc("/people_priem", people.PriemPage)
	router.HandleFunc("/people_otpusk", people.OtpuskPage)
	router.HandleFunc("/people_perevod", people.PerevodPage)

	router.HandleFunc("/control_otbor_menu", control_otbor.PageControlOtborMenu)
	router.HandleFunc("/control_otbor_term", control_otbor.PageControlOtborTerm)

	router.HandleFunc("/database", PageDatabase)
	router.HandleFunc("/otbor_refresh", OtborRefresh)
	router.HandleFunc("/otbor_index", OtborIndex)

	router.HandleFunc("/otbor_edit/{id:[0-9]+}", OtborEditPage).Methods("GET")
	router.HandleFunc("/otbor_edit/{id:[0-9]+}", OtborEditHandler).Methods("POST")
	router.HandleFunc("/otbor_delete/{id:[0-9]+}", OtborDeleteHandler)

	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./templates/home_page.html",
		"./templates/base_layout.html",
		"./templates/footer.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func OtborRefresh(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/otbor_refresh" {
		http.NotFound(w, r)
		return
	}

	db.ClearTableOtbor()
	res := db.InsertOtborFromFile() + "  |  "
	db.ClearTableTerminals()
	res += db.InsertTerminalsFromFile() + "  |  "
	db.ClearTableDepartments()
	res += db.InsertDepartmentsFromFile()

	//inf := Info{res}
	files := []string{
		"./templates/otbor/otbor_refresh.html",
		"./templates/base_layout.html",
		"./templates/footer.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, res)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func PageDatabase(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/database" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./templates/otbor/menu_db.html",
		"./templates/base_layout.html",
		"./templates/footer.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
