package main

import (
	"fmt"
	groupieGeo "groupieGeo/Functions"
	"log"
	"net/http"
	"time"
)

func init() {
	fmt.Println("starting the main function")
	before := time.Now()
	var apis groupieGeo.APIS
	err := groupieGeo.ExtractToStruct("https://groupietrackers.herokuapp.com/api", &apis)
	if err != nil {
		fmt.Printf("Error while extracting the main apis: %v\n", err)
		return
	}
	err = groupieGeo.ExtractToStruct(apis.Artists, &groupieGeo.AllArtists)
	if err != nil {
		fmt.Printf("Error while extracting the artists: %v\n", err)
		return
	}
	var locations groupieGeo.Locations
	err = groupieGeo.ExtractToStruct(apis.Locations, &locations)
	if err != nil {
		fmt.Printf("Error while extracting the locations: %v\n", err)
		return
	}
	var dates groupieGeo.Dates
	err = groupieGeo.ExtractToStruct(apis.Dates, &dates)
	if err != nil {
		fmt.Printf("Error while extracting the dates: %v\n", err)
		return
	}
	var relations groupieGeo.Relation
	err = groupieGeo.ExtractToStruct(apis.Relation, &relations)
	if err != nil {
		fmt.Printf("Error while extracting the relations: %v\n", err)
		return
	}
	for i, v := range locations.Index {
		groupieGeo.AllArtists[i].LocationsForArtist = append(groupieGeo.AllArtists[i].LocationsForArtist, v.Locations...)
		groupieGeo.AllArtists[i].DatesForArtist = append(groupieGeo.AllArtists[i].DatesForArtist, dates.Index[i].Dates...)
		groupieGeo.AllArtists[i].RelationForArtist = relations.Index[i].Relation
	}
	fmt.Println(time.Since(before))
}

func main() {
	http.HandleFunc("/", groupieGeo.HomeHandler)
	http.HandleFunc("/artists", groupieGeo.MainHandler)
	http.HandleFunc("/view", groupieGeo.ViewHandler)
	http.HandleFunc("/search", groupieGeo.SearchHandler)
	http.HandleFunc("/filtered", groupieGeo.FilterHandler)
	http.Handle("/Templates/", http.StripPrefix("/Templates/", http.FileServer(http.Dir("Templates"))))
	http.Handle("/Style/", http.StripPrefix("/Style/", http.FileServer(http.Dir("Style"))))
	groupieGeo.OpenBrowser("http://localhost:1221")
	fmt.Println("Starting server on http://localhost:1221/")
	log.Fatal(http.ListenAndServe(":1221", nil))
}
