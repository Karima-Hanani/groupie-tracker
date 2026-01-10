package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, r *http.Request, msg string, status int) {
	// w.WriteHeader(status)
	fmt.Println("Header :", w.Header())
	tmpl, err := template.ParseFiles("template/errorpage.html")
	fmt.Println("\nerror ", msg, "status ", status)
	if err != nil {
		fmt.Println("\n\nParse errorpage err :", err)
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
	err = tmpl.ExecuteTemplate(w, "errorpage.html", data)
	if err != nil {
		fmt.Println("execute errorpage err :", err)
		// http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		http.ServeFile(w, r, "/static/error.html")
		return
	}

	//

	// fmt.Fprintf(w, "Error %d: %s", status, http.StatusText(status))
}
