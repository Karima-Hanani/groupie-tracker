package handler

import (
	"net/http"
	"os"
	"path/filepath"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static/" {
		ErrorPage(w, r,"Access Forbidden", http.StatusForbidden)
		return
	}

	file := r.URL.Path[len("/static/"):]
	path := filepath.Join("static", file)
	if _, err := os.Stat(path); err != nil {
		ErrorPage(w, r,"File Not Found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, path)
}
