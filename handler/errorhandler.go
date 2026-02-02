package handler

import (
	"bytes"
	"net/http"
)

// ErrorPage renders an error page with the given message and status code.
func ErrorPage(w http.ResponseWriter, r *http.Request, msg string, status int) {
	
	data := struct {
		Message string
		Status  int
		}{
			Message: msg,
			Status:  status,
		}
		
	var buff bytes.Buffer
	err := templates.ExecuteTemplate(&buff, "errorpage.html", data)
	if err != nil {
		w.WriteHeader(500)
		templates.ExecuteTemplate(w, "error.html", data)
		return
		}
	w.WriteHeader(status)
	w.Write(buff.Bytes())
}
