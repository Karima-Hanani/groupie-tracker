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
	ID           int       `json:"id"`
	Image        string    `json:"image"`
	Name         string    `json:"name"`
	Members      []string  `json:"members"`
	CreationDate int       `json:"creationDate"`
	FirstAlbum   string    `json:"firstAlbum"`
	// Locations    Locations `json:"locations"`
	// ConcertDates Dates     `json:"concertDates"`
	// Relations    Relations `json:"relations"`
}

type Locations struct {
	ID       int      `json:"id"`
	Location []string `json:"locations"`
	// Dates    Dates    `json:"dates"`
}
type Dates struct {
	ID   int      `json:"id"`
	Date []string `json:"dates"`
}
type Relations struct {
	ID             int `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func main() {
	counter := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			http.Error(w, "failed to fetch data", http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()
		var artists []Artist
		err = json.NewDecoder(res.Body).Decode(&artists)
		if err != nil {
			fmt.Println("Naaah something's wrong dude :/\n", err)
			return
		}
		
		res1, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
		if err != nil {
			http.Error(w, "failed to fetch data", http.StatusInternalServerError)
			return
		}
		defer res1.Body.Close()
		locationWrapper:= struct{
			Index []Locations `json:"index"`
		}{}
		var locations []Locations
		err = json.NewDecoder(res1.Body).Decode(&locationWrapper)
		if err != nil {
			fmt.Println("Naaah something's wrong dude :/\n", err)
			return
		}
		locations=locationWrapper.Index

		res2, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
		if err != nil {
			http.Error(w, "failed to fetch data", http.StatusInternalServerError)
			return
		}
		defer res2.Body.Close()
		var dates []Dates
		datesWrapper:= struct{
			Index []Dates `json:"index"`
		}{}
		err = json.NewDecoder(res2.Body).Decode(&datesWrapper)
		if err != nil {
			fmt.Println("Naaah something's wrong dude :/\n", err)
			return
		}
		dates=datesWrapper.Index

		res3, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
		if err != nil {
			http.Error(w, "failed to fetch relations", http.StatusInternalServerError)
			return
		}
		defer res3.Body.Close()
		var relations []Relations
		var relationsWrapper struct{
			Index []Relations `json:"index"`
		}
		// fmt.Printf("res3 \n\n%#v\n\n", res3)
		err = json.NewDecoder(res3.Body).Decode(&relationsWrapper)
		if err != nil {
			fmt.Println("Naaah! something's wrong with relations dude :/\n", err)
			return
		}
		relations=relationsWrapper.Index

		fmt.Printf("Artists====== \n\n%#v\n\n", artists[0])
		fmt.Printf("Locations====== \n\n%#v\n\n", locations[0])
		fmt.Printf("Dates====== \n\n%#v\n\n", dates[0])
		fmt.Printf("Relations====== \n\n%#v\n\n", relations[0])
		counter++
		fmt.Println(counter)
	})
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
	
}