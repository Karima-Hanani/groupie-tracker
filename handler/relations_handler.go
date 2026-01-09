package handler

import (
	"net/http"
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
	
}
