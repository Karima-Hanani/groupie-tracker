package handler

import (
	"bytes"
	"html/template"
	"net/http"

	"groupie-tracker/fetcher"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// check for path
	if r.URL.Path != "/" {
		ErrorPage(w, r, "Page not found.", 404)
		return
	}

	// check for method
	if r.Method != "GET" {
		ErrorPage(w, r, "Method is not supported.", 405)
		return
	}

	artists, err := fetcher.FetchArtists()
	if err != nil {
		ErrorPage(w, r, "Failed to load artists.", 500)
		return
	}

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		ErrorPage(w, r, "Failed to load template.", 500)
		return
	}

	var buff bytes.Buffer
	// Render the homepage template with the fetched data
	err = tmpl.ExecuteTemplate(&buff, "index.html", artists)
	if err != nil {
		ErrorPage(w, r, "Failed to render template.", 500)
		return
	}
	w.Write(buff.Bytes())
}
