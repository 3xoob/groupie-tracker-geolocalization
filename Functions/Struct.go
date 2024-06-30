package groupieGeo

type APIS struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Artists struct {
	ID                 int      `json:"id"`
	Image              string   `json:"image"`
	Name               string   `json:"name"`
	Members            []string `json:"members"`
	CreationDate       int      `json:"creationDate"`
	FirstAlbum         string   `json:"firstAlbum"`
	LocationsForArtist []string
	DatesForArtist     []string
	RelationForArtist  map[string][]string
	Coordinates        []Cords
}

type Locations struct {
	Index []struct {
		Locations []string `json:"locations"`
	} `json:"index"`
}

type Dates struct {
	Index []struct {
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Relation struct {
	Index []struct {
		Relation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type ArtistPageData struct {
	Artists            []Artists
	SearchLocations    []string
	SearchCreationDate []int
	SearchFirstAlbum   []string
	NumberOfMember     [8]string
	Filter             Filters
}

type Filters struct {
	CreationDateRange [2]int
	FirstAlbumRange   [2]string
	SelectedLocations string
	SelectedMembers   []string
}

type Cords struct {
	Name string
	Lat  float64
	Lng  float64
}

var AllArtists []Artists
