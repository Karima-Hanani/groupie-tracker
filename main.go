package main

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/handler"
)


func main() {
	templates, err := template.ParseFiles(
		"templates/errorpage.html",
		"templates/index.html",
		"templates/details.html",
		"static/error.html",
	)
	
	if err != nil {
		fmt.Println("Error parsing templates",err)
		return 
	}
	
	handler.Temp(templates)

	http.HandleFunc("/", handler.HomePage)
	http.HandleFunc("/static/", handler.StaticHandler)
	http.HandleFunc("/details", handler.DetailsHandler)
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
