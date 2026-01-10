# Groupie-Tracker

## Description

Groupie-Tracker is a web application that displays information about music artists, their members, concert locations, and dates. The backend is written in Go, and the frontend uses HTML templates and CSS for rendering. The application fetches data from the [Groupie Tracker API](https://groupietrackers.herokuapp.com/api), allowing users to explore artists and their concert history.

## What is API ? 
An API (Application Programming Interface) is a set of rules that allows one software application to communicate with another. In this project, we use the Groupie Tracker API to retrieve artist and concert information.


## Features

* Web interface to browse all artists.
* View artist details: name, image, creation date, first album, and members.
* Display concert locations and dates.
* Show relation between locations and concert dates.
* Dynamic rendering using HTML templates.
* Error handling for invalid input or server errors.

## How to Run

Clone the repository:

```bash
git clone <your-repo-link>
cd groupie-tracker
```

Start the server locally:

```bash
go run .
# then open http://localhost:8080
```

## Implementation Details

### Backend (Go)

* Uses Go's `net/http` package for the web server.
* Routes:

  * `GET /` — serves the home page with a list of artists.
  * `GET /details?id=<id>` — displays detailed artist information including locations, dates, and relations.
* Fetches data from the Groupie Tracker API using HTTP requests.

### Frontend (HTML & CSS)

* Templates live in the `templates/` directory and use `html/template`.
* CSS files are in the `static/` directory.
* Artist page displays multiple data points and uses `range` loops to iterate over slices and maps.
* Relation between locations and dates can be displayed as a `<select>` dropdown or as lists.


## Project Structure

```bash
.
├── main.go            # Entry point for the server
├── server/            # HTTP handlers and routing logic
├── templates/         # HTML template files (index.html, artist.html)
├── static/            # CSS files and images
└── README.md
```

## Authors

- <span style="color:pink; font-family: Verdana; font-weight: bold; font-size:16;">Karima Hanani</span>
- <span style="color:pink; font-family: Verdana; font-weight: bold; font-size:16;">Meryam Aadlani</span>
