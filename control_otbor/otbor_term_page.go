package control_otbor

import (
	"html/template"
	"log"
	"net/http"
)

func PageControlOtborTerm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/control_otbor_term" {
		http.NotFound(w, r)
		return
	}

	res := OtborTextTerm()

	//inf := Info{res}
	files := []string{
		"./templates/control_otbor_menu.html",
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
