// electrum project main.go
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Db *sql.DB

func main() {

	Db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		//return err
		panic(err)
	}
	defer Db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/database", PageDatabase)
	router.HandleFunc("/otbor_refresh", OtborRefresh)
	router.HandleFunc("/otbor_index", OtborIndex)

	router.HandleFunc("/otbor_edit/{id:[0-9]+}", OtborEditPage).Methods("GET")
	router.HandleFunc("/otbor_edit/{id:[0-9]+}", OtborEditHandler).Methods("POST")
	router.HandleFunc("/otbor_delete/{id:[0-9]+}", OtborDeleteHandler)

	//t, _ := template.ParseFiles("templates/index-go.html", "templates/base-go.html")
	//name := "world"
	//t.Execute(os.Stdout, name)

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

	ClearTableOtbor()
	res := InsertOtborFromFile() + "  |  "
	ClearTableTerminals()
	res += InsertTerminalsFromFile() + "  |  "
	ClearTableDepartments()
	res += InsertDepartmentsFromFile()

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
