package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/fetcher"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPage(w, r, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	showID := r.URL.Query().Get("id")
	fmt.Println("showID:", showID)
	if showID == "" {
		ErrorPage(w, r, "ID is required.", http.StatusBadRequest)
		return
	}

	Artist, err := fetcher.FetchArtist(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load artist.", http.StatusInternalServerError)
		return
	}

	Dates, err := fetcher.FetchDates(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load dates.", http.StatusInternalServerError)
		return
	}

	Locations, err := fetcher.FetchLocations(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load locations.", http.StatusInternalServerError)
		return
	}

	Relation, err := fetcher.FetchRelations(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load relations.", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("template/details.html")
	if err != nil {
		fmt.Println("parse details err :", err)
		ErrorPage(w, r, "Failed to load template.", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "details.html", map[string]any{
		"Relations": Relation,
		"Artist":    Artist,
		"Dates":     Dates,
		"Locations": Locations,
	})
	if err != nil {
		fmt.Println("details err :", err)
		ErrorPage(w, r, "Failed to render template.", http.StatusInternalServerError)
		return
	}
}
