package groupieGeo

import (
	"strconv"
	"strings"
)

func ApplyFilters(artists []Artists, filter Filters) []Artists {
	filteredArtists := artists
	var tempFilter []Artists
	minCreationDate := strconv.Itoa(filter.CreationDateRange[0])
	maxCreationDate := strconv.Itoa(filter.CreationDateRange[1])
	minFirstAlbum := filter.FirstAlbumRange[0]
	maxFirstAlbum := filter.FirstAlbumRange[1]

	if len(filter.SelectedMembers) > 0 {
		tempFilter = []Artists{}
		for _, artist := range artists {
			for _, num := range filter.SelectedMembers {
				numInt, _ := strconv.Atoi(num)
				if len(artist.Members) == numInt {
					tempFilter = append(tempFilter, artist)
				}
			}
		}
		filteredArtists = tempFilter
	}

	if minCreationDate != "" || maxCreationDate != "" {
		if minCreationDate != "1950" || maxCreationDate != "2024" {
			tempFilter = []Artists{}
			for _, artist := range filteredArtists {
				if (minCreationDate == "" || strconv.Itoa(artist.CreationDate) >= minCreationDate) &&
					(maxCreationDate == "" || strconv.Itoa(artist.CreationDate) <= maxCreationDate) {
					tempFilter = append(tempFilter, artist)
				}
			}
			filteredArtists = tempFilter
		}
	}

	if minFirstAlbum != "" || maxFirstAlbum != "" {
		if minFirstAlbum != "1950" || maxFirstAlbum != "2024" {
			tempFilter = []Artists{}

			for _, artist := range filteredArtists {
				year := strings.Split(artist.FirstAlbum, "-")[2]
				if (minFirstAlbum == "" || year >= minFirstAlbum) &&
					(maxFirstAlbum == "" || year <= maxFirstAlbum) {
					tempFilter = append(tempFilter, artist)
				}
			}
			filteredArtists = tempFilter
		}
	}

	if filter.SelectedLocations != "" {
		tempFilter = []Artists{}
		for _, artist := range filteredArtists {
			for _, va := range artist.LocationsForArtist {
				if va == filter.SelectedLocations {
					tempFilter = append(tempFilter, artist)
				}
			}
		}
		filteredArtists = tempFilter
	}

	return RemoveDuplicates(filteredArtists)
}

func RemoveDuplicates(arr []Artists) []Artists {
	encountered := map[int]bool{}
	result := []Artists{}

	for _, v := range arr {
		if !encountered[v.ID] {
			encountered[v.ID] = true
			result = append(result, v)
		}
	}

	return result
}
