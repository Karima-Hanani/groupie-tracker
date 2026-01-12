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
}

type Locations struct {
	ID       int      `json:"id"`
	Location []string `json:"locations"`
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
		return nil, fmt.Errorf("failed to fetch artist %#v", err)
	}
	defer res.Body.Close()

	var artists []Artist
	
	err = json.NewDecoder(res.Body).Decode(&artists)
	
	artists[20].Image = "static/forbiden.png"
	if err != nil {
		return nil, fmt.Errorf("failed to decode artists %#v", err)
	}
	return artists, nil
}

func FetchArtist(id string) (Artist, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + fmt.Sprint(id))
	if err != nil {
		return Artist{}, fmt.Errorf("failed to fetch artist %#v", err)
	}
	defer res.Body.Close()

	var artists Artist

	err = json.NewDecoder(res.Body).Decode(&artists)
	if id == "20" {
		artists.Image = "static/forbiden.png"
	}
	if err != nil {
		return Artist{}, fmt.Errorf("failed to decode artist %#v", err)
	}
	return artists, nil
}

func FetchLocations(id string) (Locations, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + fmt.Sprint(id))
	if err != nil {
		return Locations{}, fmt.Errorf("failed to fetch locations %#v", err)
	}
	defer res.Body.Close()

	var locations Locations

	err = json.NewDecoder(res.Body).Decode(&locations)
	if err != nil {
		return Locations{}, fmt.Errorf("failed to decode locations %#v", err)
	}

	return locations, nil
}

func FetchDates(id string) (Dates, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + fmt.Sprint(id))
	if err != nil {
		return Dates{}, fmt.Errorf("failed to fetch dates %#v", err)
	}
	defer res.Body.Close()

	var dates Dates

	err = json.NewDecoder(res.Body).Decode(&dates)
	if err != nil {
		return Dates{}, fmt.Errorf("failed to decode dates %#v", err)
	}

	return dates, nil
}

func FetchRelations(id string) (Relations, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + fmt.Sprint(id))
	if err != nil {
		return Relations{}, fmt.Errorf("failed to fetch relations %#v", err)
	}
	defer res.Body.Close()

	var relation Relations
	
	err = json.NewDecoder(res.Body).Decode(&relation)
	if err != nil {
		return Relations{}, fmt.Errorf("failed to decode relations %#v", err)
	}
	return relation, nil
}
