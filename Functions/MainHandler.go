package groupieGeo

import (
	"fmt"
	"net/http"
	"text/template"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists" {
		// error 404
		ErrorHandler(w, r, http.StatusNotFound, "")
		return
	}
	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		// error 500
		ErrorHandler(w, r, http.StatusInternalServerError, "")
		fmt.Printf("err: %v\n", err)
		return
	}
	Flocation, Falbum, Fcreation := Filter(AllArtists)

	filter := ParseFilters(r)

	AllData := ArtistPageData{
		Artists:            AllArtists,
		SearchLocations:    Flocation,
		SearchCreationDate: Fcreation,
		SearchFirstAlbum:   Falbum,
		NumberOfMember:     [8]string{"1", "2", "3", "4", "5", "6", "7", "8"},
		Filter:             filter,
	}

	tmpl.Execute(w, AllData)
}
