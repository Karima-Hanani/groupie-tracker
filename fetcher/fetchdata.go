package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
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
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func FetchArtists() ([]Artist, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		// http.Error(w, "failed to fetch data", http.StatusInternalServerError)
		return nil, fmt.Errorf("failed to fetch data %#v", err)
	}
	defer res.Body.Close()

	var artists []Artist

	err = json.NewDecoder(res.Body).Decode(&artists)
	if err != nil {
		// fmt.Println("Naaah something's wrong dude :/\n", err)
		return nil, fmt.Errorf("failed to decode data %#v", err)
	}
	return artists, nil
}

func FetchLocations() ([]Locations, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data %#v", err)
	}
	defer res.Body.Close()
	locationWrapper := struct {
		Index []Locations `json:"index"`
	}{}
	var locations []Locations
	err = json.NewDecoder(res.Body).Decode(&locationWrapper)
	if err != nil {
		// fmt.Println("Naaah something's wrong dude :/\n", err)
		return nil, fmt.Errorf("failed to decode data %#v", err)
	}
	locations = locationWrapper.Index
	return locations, nil
}


func FetchDates() ([]Dates, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data %#v", err)
	}
	defer res.Body.Close()
	var dates []Dates
	datesWrapper := struct {
		Index []Dates `json:"index"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&datesWrapper)
	if err != nil {
		// fmt.Println("Naaah something's wrong dude :/\n", err)
		return nil, fmt.Errorf("failed to decode data %#v", err)
	}
	dates = datesWrapper.Index
	return dates, nil
}
func FetchRelations() ([]Relations, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch relations %#v", err)
	}
	defer res.Body.Close()
	var relations []Relations
	var relationsWrapper struct {
		Index []Relations `json:"index"`
	}
	err = json.NewDecoder(res.Body).Decode(&relationsWrapper)
	if err != nil {
		// fmt.Println("Naaah! something's wrong with relations dude :/\n", err)
		return nil, fmt.Errorf("failed to decode data %#v", err)
	}
	relations = relationsWrapper.Index
	return relations, nil
}