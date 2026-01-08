package handler

import (
	"net/http"

	"groupie-tracker/fetcher"
	"html/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// check for path
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	// check for method
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	artists, err := fetcher.FetchArtists()
	if err != nil {
		http.Error(w, "Failed to load artists.", http.StatusInternalServerError)
		return
	}
	locations, err := fetcher.FetchLocations()
	if err != nil {
		http.Error(w, "Failed to load locations.", http.StatusInternalServerError)
		return
	}
	dates, err := fetcher.FetchDates()
	if err != nil {
		http.Error(w, "Failed to load dates.", http.StatusInternalServerError)
		return
	}
	relations, err := fetcher.FetchRelations()
	if err != nil {
		http.Error(w, "Failed to load relations.", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("handler/homepage.html")
	if err != nil {
		http.Error(w, "Failed to load template.", http.StatusInternalServerError)
		return
	}
	// Render the homepage template with the fetched data
	err = tmpl.ExecuteTemplate(w, "homepage.html", map[string]interface{}{
		"Artists":   artists,
		"Locations": locations,
		"Dates":     dates,
		"Relations": relations,
	})
	if err != nil {
		http.Error(w, "Failed to render template.", http.StatusInternalServerError)
	}
}
