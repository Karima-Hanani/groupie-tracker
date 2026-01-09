package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("http://localhost:8080/")
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println("error")
		return
	}
}
