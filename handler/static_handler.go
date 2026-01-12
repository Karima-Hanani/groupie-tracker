package handler

import (
	"net/http"
	"os"
	"path/filepath"
)

// StaticHandler serves static files while preventing directory access.
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static/" {
		ErrorPage(w, r, "Access Forbidden", 403)
		return
	}

	file := r.URL.Path[len("/static/"):]
	path := filepath.Join("static", file)
	if _, err := os.Stat(path); err != nil {
		ErrorPage(w, r, "File Not Found", 404)
		return
	}
	http.ServeFile(w, r, path)
}
