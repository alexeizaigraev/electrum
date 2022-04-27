package people

import (
	"html/template"
	"log"
	"net/http"
)

func PostAllPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/people_postall" {
		http.NotFound(w, r)
		return
	}

	res := PostAllMain()

	//inf := Info{res}
	files := []string{
		"./templates/out5.html",
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
