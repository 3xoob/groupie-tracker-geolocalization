package groupieGeo

import (
	"net/http"
	"strconv"
)

func ParseFilters(r *http.Request) Filters {
	var filter Filters

	startYear, _ := strconv.Atoi(r.FormValue("MinCreationDate"))
	endYear, _ := strconv.Atoi(r.FormValue("MaxCreationDate"))
	filter.CreationDateRange = [2]int{startYear, endYear}
	startAlbum := r.FormValue("MinFirstAlbum")
	endAlbum := r.FormValue("MaxFirstAlbum")
	filter.FirstAlbumRange = [2]string{startAlbum, endAlbum}
	filter.SelectedLocations = r.FormValue("locations")
	filter.SelectedMembers = r.Form["NumMembers"]

	return filter
}
