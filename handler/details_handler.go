package handler

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"

	"groupie-tracker/fetcher"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPage(w, r, "Method is not supported.", 405)
		return
	}

	showID := r.URL.Query().Get("id")
	if showID == "" {
		ErrorPage(w, r, "ID is required.", 400)
		return
	}

	Artist, err := fetcher.FetchArtist(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load artist.", 500)
		return
	}

	if showID == "21" {
		Artist.Image = "static/forbiden.png"
	}

	Dates, err := fetcher.FetchDates(showID)
	if err != nil {
		ErrorPage(w, r, "Failed to load dates.", 500)
		return
	}

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
		ErrorPage(w, r, "Failed to load template.", 500)
		return
	}

	var buff bytes.Buffer
	err = tmpl.ExecuteTemplate(&buff, "details.html", map[string]any{
		"Relations": Relation,
		"Artist":    Artist,
		"Dates":     Dates,
		"Locations": Locations,
	})
	if err != nil {
		ErrorPage(w, r, "Failed to render template.", 500)
		return
	}
	w.Write(buff.Bytes())
}
