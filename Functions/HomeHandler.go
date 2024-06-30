package groupieGeo

import (
	"fmt"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "")
		return
	}

	if r.URL.Path != "/" {
		// error 404
		ErrorHandler(w, r, http.StatusNotFound, "")
		return
	}

	tmpl, err := template.ParseFiles("Templates/home.html")
	if err != nil {
		// error 500
		ErrorHandler(w, r, http.StatusInternalServerError, "")
		fmt.Printf("err: %v\n", err)
		return
	}

	tmpl.Execute(w, nil)
}
