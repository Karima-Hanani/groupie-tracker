package handler

import (
	"html/template"
	"net/http"

	"groupie-tracker/fetcher"
)

func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != "GET" {
		ErrorPage(w, r, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	showID := r.URL.Query().Get("show")
	if showID == "" {
		ErrorPage(w, r, "ID is required.", http.StatusBadRequest)
		return
	}

	Relation, err := fetcher.FetchRelations(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load relations.", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("handler/relations.html")
	if err != nil {
		ErrorPage(w, r, "Failed to load template.", http.StatusInternalServerError)
		return
	}

	err=tmpl.ExecuteTemplate(w, "relations.html", Relation)
	if err != nil {
		ErrorPage(w, r, "Failed to render template.", http.StatusInternalServerError)
		return
	}
}