package handler

import (
	"bytes"
	"html/template"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, r *http.Request, msg string, status int) {
	w.WriteHeader(status)
	tmpl, err := template.ParseFiles("template/errorpage.html")
	if err != nil {
		http.ServeFile(w, r, "/static/error.html")
		return
	}

	data := struct {
		Message string
		Status  int
	}{
		Message: msg,
		Status:  status,
	}
	var buff bytes.Buffer
	err = tmpl.ExecuteTemplate(&buff, "errorpage.html", data)
	if err != nil {
		http.ServeFile(w, r, "/static/error.html")
		return
	}
	w.Write(buff.Bytes())
}
