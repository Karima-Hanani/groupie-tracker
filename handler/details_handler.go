package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"groupie-tracker/fetcher"
)

// DetailsHandler handles the details page for a specific artist.
func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPage(w, r, "Method is not supported.", 405)
		return
	}
	showID := r.URL.Query().Get("id")
	if showID == "" {
		ErrorPage(w, r, "Page Not Found.", 404)
		return
	}

	Artist, err := fetcher.FetchArtist(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load artist.", 500)
		return
	}
	if Artist.ID == 0 {
		ErrorPage(w, r, "Artist Not found", 404)
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
		ErrorPage(w, r, "Failed to loaddd locations.", 500)
		return
	}

	Relation, err := fetcher.FetchRelations(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load relations.", 500)
		return
	}

	// Render the template with the fetched data
	var buff bytes.Buffer
	err = templates.ExecuteTemplate(&buff, "details.html", map[string]any{
		"Relations": Relation,
		"Artist":    Artist,
		"Dates":     Dates,
		"Locations": Locations,
	})
	if err != nil {
		fmt.Println("Render : ", err)
		ErrorPage(w, r, "Failed to render template.", 500)
		return
	}
	w.Write(buff.Bytes())
}
