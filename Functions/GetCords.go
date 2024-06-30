package groupieGeo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type GeocodingResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
	Status string `json:"status"`
}

func GetCords(location string) (float64, float64, error) {
	apiKey := "AIzaSyACt5oDURFW4CJ0iI_cNKOaYJXXf9vu8vU"

	location = strings.ReplaceAll(location, "-", ", ")
	location = strings.ReplaceAll(location, "_", " ")
	location = url.QueryEscape(location)

	apiURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", location, apiKey)

	apiData, err := GetData(apiURL)
	if err != nil {
		return 0.0, 0.0, fmt.Errorf("failed to retrieve data from API: %w", err)
	}

	var response GeocodingResponse

	err = json.Unmarshal(apiData, &response)
	if err != nil {
		return 0.0, 0.0, fmt.Errorf("failed to unmarshal API response: %w", err)
	}

	if response.Status != "OK" {
		return 0.0, 0.0, fmt.Errorf("API responded with status: %s", response.Status)
	}

	if len(response.Results) == 0 {
		return 0.0, 0.0, fmt.Errorf("no results found for the location: %s", location)
	}

	lat := response.Results[0].Geometry.Location.Lat
	lng := response.Results[0].Geometry.Location.Lng

	return lat, lng, nil
}
