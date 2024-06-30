package groupieGeo

func Filter(artists []Artists) ([]string, []string, []int) {
	uniqueSearchLocations := make(map[string]struct{})
	uniqueFirstAlbums := make(map[string]struct{})
	uniqueCreationDates := make(map[int]struct{})

	for _, artist := range artists {
		if _, found := uniqueFirstAlbums[artist.FirstAlbum]; !found {
			uniqueFirstAlbums[artist.FirstAlbum] = struct{}{}
		}
		for _, location := range artist.LocationsForArtist {
			if _, found := uniqueSearchLocations[location]; !found {
				uniqueSearchLocations[location] = struct{}{}
			}
		}

		if _, found := uniqueCreationDates[artist.CreationDate]; !found {
			uniqueCreationDates[artist.CreationDate] = struct{}{}
		}
	}

	var searchLocs []string
	for loc := range uniqueSearchLocations {
		searchLocs = append(searchLocs, loc)
	}

	var firstAlbums []string
	for album := range uniqueFirstAlbums {
		firstAlbums = append(firstAlbums, album)
	}

	var creationDates []int
	for date := range uniqueCreationDates {
		creationDates = append(creationDates, date)
	}

	return searchLocs, firstAlbums, creationDates
}
