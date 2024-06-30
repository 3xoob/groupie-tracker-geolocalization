package groupieGeo

import (
	"fmt"
	"net/http"
	"text/template"
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "")
		return
	}

	tmpl, err := template.ParseFiles("Templates/filter.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "")
		fmt.Printf("err: %v\n", err)
		return
	}
	Flocation, Falbum, Fcreation := Filter(AllArtists)

	filter := ParseFilters(r)

	filteredArtists := ApplyFilters(AllArtists, filter)

	AllData := ArtistPageData{
		Artists:            filteredArtists,
		SearchLocations:    Flocation,
		SearchCreationDate: Fcreation,
		SearchFirstAlbum:   Falbum,
		NumberOfMember:     [8]string{"1", "2", "3", "4", "5", "6", "7", "8"},
		Filter:             filter,
	}

	tmpl.Execute(w, AllData)
}
