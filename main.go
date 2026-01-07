package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	resp2, err2 := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}
	defer resp2.Body.Close()
	defer resp.Body.Close()

	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		panic(err)
	}

	var todo Todo
	err= json.NewDecoder(resp2.Body).Decode(&todo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %d, Title: %s\n\n",todo.ID,todo.Title)
	fmt.Println("=================== Users ==================")

	for _, u := range users {
		fmt.Println(u.Name, "-", u.Email)
	}
}
