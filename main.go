package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type User struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }
// type Todo struct{
// 	ID int `json:"id"`
// 	Title string `json:"title"`
// }

// func main() {
// 	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
// 	resp2, err2 := http.Get("https://jsonplaceholder.typicode.com/todos/1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err2 != nil {
// 		panic(err2)
// 	}
// 	defer resp2.Body.Close()
// 	defer resp.Body.Close()

// 	var users []User
// 	err = json.NewDecoder(resp.Body).Decode(&users)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var todo Todo
// 	err= json.NewDecoder(resp2.Body).Decode(&todo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("ID: %d, Title: %s\n\n",todo.ID,todo.Title)
// 	fmt.Println("=================== Users ==================")

// 	for _, u := range users {
// 		fmt.Println(u.Name, "-", u.Email)
// 	}
// }

type Artist struct {
	ID int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
	Locations Locations `json:"locations"`
	ConcertDates Dates `json:"concertDates"`
	Relations Relations `json:"relations"`
}

type Locations struct {
	ID int `json:"id"`
	Location []string `json:"locations"`
	Dates Dates `json:"dates"`
}
type Dates struct {
	ID int `json:"id"`
	Date []string `json:"dates"`
}
type Relations struct {
	ID int `json:"id"`
	DatesLocations map[string]string
}

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		res,err :=http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			http.Error(w,"failed to fetch data",http.StatusInternalServerError)
			return	
		}
		defer res.Body.Close()
		var artists []Artist

		json.NewDecoder(res.Body).Decode(&artists)
		res,err =http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			http.Error(w,"failed to fetch data",http.StatusInternalServerError)
			return	
		}
		defer res.Body.Close()
		var locations []Locations
		json.NewDecoder(res.Body).Decode(&locations)

		res,err =http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			http.Error(w,"failed to fetch data",http.StatusInternalServerError)
			return	
		}
		defer res.Body.Close()
		var dates []Dates
		json.NewDecoder(res.Body).Decode(&dates)

		res,err =http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			http.Error(w,"failed to fetch data",http.StatusInternalServerError)
			return	
		}
		defer res.Body.Close()
		var relations []Dates
		json.NewDecoder(res.Body).Decode(&relations)

		fmt.Printf("Artists %#v\n\n", artists[1])
		fmt.Printf("Locations %#v\n\n", locations[1])
		fmt.Printf("Dates %#v\n\n", dates[1])
		fmt.Printf("Relations %#v\n\n", relations[1])
	})
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080",nil)
}
