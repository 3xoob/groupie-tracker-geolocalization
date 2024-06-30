package groupieGeo

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	searchQuery := r.FormValue("theArtist")
	arr := []string{
		" -Location", " -Member", " -First Album", " -Creation Date", " - Location", " - Member", " - First Album",
		" - Creation Date", " -location", " -member", " -first album", " -creation date", " -firstalbum", " -creationdate", " -FirstAlbum",
		" -CreationDate", " - location", " - member", " - first album", " - creation date", "-Location", "-Member", "-First Album",
		"-Creation Date", "- Location", "- Member", "- First Album", "- Creation Date", "-location", "-member", "-first album",
		"-creation date", "-firstalbum", "-creationdate", "-FirstAlbum", "-CreationDate", "- location", "- member", "- first album", "- creation date",
		" -artist/band", " - artist/band", "-artist/band", "- artist/band",
		" -Artist/band", " - Artist/band", "-Artist/band", "- Artist/band",
		" -artist/Band", " - artist/Band", "-artist/Band", "- artist/Band",
		" -Artist/Band", " - Artist/Band", "-Artist/Band", "- Artist/Band",
		" -artist/ band", " - artist/ band", "-artist/ band", "- artist/ band",
		" -Artist/ band", " - Artist/ band", "-Artist/ band", "- Artist/ band",
		" -artist/ Band", " - artist/ Band", "-artist/ Band", "- artist/ Band",
		" -Artist/ Band", " - Artist/ Band", "-Artist/ Band", "- Artist/ Band",
		" -artist / band", " - artist / band", "-artist / band", "- artist / band",
		" -Artist / band", " - Artist / band", "-Artist / band", "- Artist / band",
		" -artist / Band", " - artist / Band", "-artist / Band", "- artist / Band",
		" -Artist / Band", " - Artist / Band", "-Artist / Band", "- Artist / Band",
		" -artist /band", " - artist /band", "-artist /band", "- artist /band",
		" -Artist /band", " - Artist /band", "-Artist /band", "- Artist /band",
		" -artist /Band", " - artist /Band", "-artist /Band", "- artist /Band",
		" -Artist /Band", " - Artist /Band", "-Artist /Band", "- Artist /Band"}

	for _, v := range arr {
		if strings.HasSuffix(searchQuery, v) {
			searchQuery = strings.TrimSuffix(searchQuery, v)
		}
	}

	var results []Artists
	uniqueArtists := make(map[int]bool)

	for _, artist := range AllArtists {
		if uniqueArtists[artist.ID] {
			continue
		}
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(searchQuery)) ||
			strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(searchQuery)) ||
			strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), strings.ToLower(searchQuery)) {
			results = append(results, artist)
			uniqueArtists[artist.ID] = true
		}
		for _, member := range artist.Members {
			if uniqueArtists[artist.ID] {
				break
			}
			if strings.Contains(strings.ToLower(member), strings.ToLower(searchQuery)) {
				results = append(results, artist)
				uniqueArtists[artist.ID] = true
				break
			}
		}
		for _, location := range artist.LocationsForArtist {
			if uniqueArtists[artist.ID] {
				break
			}
			if strings.Contains(strings.ToLower(location), strings.ToLower(searchQuery)) {
				results = append(results, artist)
				uniqueArtists[artist.ID] = true
				break
			}
		}
	}

	tmpl, err := template.ParseFiles("Templates/search.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "")
		fmt.Printf("err: %v\n", err)
		return
	}
	Flocation, Falbum, Fcreation := Filter(results)

	AllData := ArtistPageData{
		Artists:            results,
		SearchLocations:    Flocation,
		SearchCreationDate: Fcreation,
		SearchFirstAlbum:   Falbum,
	}

	tmpl.Execute(w, AllData)
}
