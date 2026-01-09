package handler

import (
	"html/template"
	"net/http"
)

func ErrorPage(w http.ResponseWriter,r *http.Request, msg string, status int) {
	w.WriteHeader(status)
	tmpl, err := template.ParseFiles("handler/errorpage.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := struct {
		Message string
		Status  int
	}{
		Message: msg,
		Status:  status,
	}
	err = tmpl.ExecuteTemplate(w, "errorpage.html", data)
	if err != nil {
		// http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		http.ServeFile(w,r,"/static/error.html")
		return
	}
	
	//

	// fmt.Fprintf(w, "Error %d: %s", status, http.StatusText(status))
}
