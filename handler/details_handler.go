package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/fetcher"
)

// DetailsHandler handles the details page for a specific artist.
func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPage(w, r, "Method is not supported.", 405)
		return
	}
	// Get artist ID from query parameters
	showID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(showID)
	if showID == "" || err != nil || (id < 1 || id > 52) {
		ErrorPage(w, r, "Bad Request.", 400)
		return
	}

	Artist, err := fetcher.FetchArtist(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load artist.", 500)
		return
	}

	Dates, err := fetcher.FetchDates(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load dates.", 500)
		return
	}
	// Clean up date strings by removing leading asterisks
	for i, v := range Dates.Date {
		if strings.HasPrefix(v, "*") {
			Dates.Date[i] = v[1:]
		}
	}

	Locations, err := fetcher.FetchLocations(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load locations.", 500)
		return
	}

	Relation, err := fetcher.FetchRelations(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load relations.", 500)
		return
	}

	tmpl, err := template.ParseFiles("template/details.html")
	if err != nil {
		fmt.Println("Error", err)
		ErrorPage(w, r, "Failed to load template.", 500)
		return
	}

	// Render the template with the fetched data
	var buff bytes.Buffer
	err = tmpl.ExecuteTemplate(&buff, "details.html", map[string]any{
		"Relations": Relation,
		"Artist":    Artist,
		"Dates":     Dates,
		"Locations": Locations,
	})
	if err != nil {
		fmt.Println("Render : ",err)
		ErrorPage(w, r, "Failed to render template.", 500)
		return
	}
	w.Write(buff.Bytes())
}
