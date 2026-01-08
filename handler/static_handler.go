package handler

import "net/http"

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static/" {
		ErrorPage(w,"Access Forbidden",http.StatusForbidden)
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
