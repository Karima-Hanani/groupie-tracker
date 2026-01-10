package main

import (
	"fmt"
	"net/http"

	"groupie-tracker/handler"
)

func main() {
	http.HandleFunc("/", handler.HomePage)
	http.HandleFunc("/static/", handler.StaticHandler)
	http.HandleFunc("/details", handler.DetailsHandler)
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
