package groupieGeo

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		errorMsg := fmt.Sprintf("Invalid artist ID: %v", err)
		ErrorHandler(w, r, http.StatusBadRequest, errorMsg)
		log.Println(errorMsg)
		return
	}
	if num < 1 || num > len(AllArtists) {
		errorMsg := fmt.Sprintf("Artist ID %d out of range", num)
		ErrorHandler(w, r, http.StatusBadRequest, errorMsg)
		log.Println(errorMsg)
		return
	}

	var cords []Cords
	for _, loc := range AllArtists[num-1].LocationsForArtist {
		lat, lng, err := GetCords(loc)
		if err != nil {
			errorMsg := fmt.Sprintf("Failed to get coordinates for location '%s': %v", loc, err)
			ErrorHandler(w, r, http.StatusInternalServerError, errorMsg)
			log.Println(errorMsg)
			return
		}
		cords = append(cords, Cords{Name: loc, Lat: lat, Lng: lng})
	}
	AllArtists[num-1].Coordinates = cords

	tmpl, err := template.ParseFiles("Templates/view.html")
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to parse template 'view.html': %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError, errorMsg)
		log.Println(errorMsg)
		return
	}

	err = tmpl.Execute(w, AllArtists[num-1])
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to execute template: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError, errorMsg)
		log.Println(errorMsg)
	}
}
