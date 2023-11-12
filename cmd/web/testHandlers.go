package main

import (
	"html/template"
	"net/http"
)

// http://localhost:4000/sb
// curl -i -X POST http://localhost:4000/sb
func (app *application) sb(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	files := []string{
		"./ui/html/pages/secondBreakfast.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		// app.logger.Error(err.Error())
		// http.Error(w, "secondBreakfast page not found", http.StatusInternalServerError)
		return
	}

	// "sb" IS NOT THE FILE NAME
	// it's the name you defined inside the template? {{define "sb"}}
	err = ts.ExecuteTemplate(w, "sb", nil)
	if err != nil {
		app.serverError(w, r, err)
		// app.logger.Error(err.Error())
		// http.Error(w, "'sb' template not found", http.StatusInternalServerError)
		return
	}

}
